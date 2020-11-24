
# %% Test
from airflow.models import DAG
from airflow.operators.dummy_operator import DummyOperator
from airflow.operators.python_operator import PythonOperator
from common.aws.s3 import S3
from airflow.models import Variable
import boto3
s3 = boto3.client('s3')

bucket = 'cdl-bi-data-science'
lp_key = 'loss_prevention/'

response = s3.list_objects_v2(
    Bucket=bucket,
    Delimiter="/",
    Prefix=lp_key
)
max_key = response["Contents"][0]["Key"]
max_last_modified = response["Contents"][0]["LastModified"]

for obj in response["Contents"]:
    if obj["LastModified"] > max_last_modified:
        max_key = obj["Key"]
        max_last_modified = obj["LastModified"]

# %%


def check_date(**kwargs):
    bucket = 'cdl-bi-data-science'
    lp_key = 'loss_prevention/'
    s3_conn = S3('s3_conn')

    recent_file = s3_conn.get_last_modified_file(
        s3_key=lp_key,
        bucket_name=bucket
    )
    print(recent_file)


def_args = {
    "owner": "QA",
    "start_date": "2020-09-04",
    "email_on_failure": False,
    "email": "robert.inkpen@compassdigital.io",
    'task_memory': 1024
}

with DAG(
    dag_id="QA_main",
    description="",  # noqa
    default_args=def_args,
    max_active_runs=1,
    catchup=False,
    schedule_interval=None,
    tags=[""]

) as dag:
    start = DummyOperator(
        task_id="start_script",
    )
    run = PythonOperator(
        task_id="QA_new_method",
        python_callable=check_date,
        provide_context=True,
        op_kwargs={

        }

    )
    start >> run

# %%
from airflow.operators.python_operator import PythonOperator
from common.aws.s3 import S3
from airflow.models import Variable

bash = '''
foo bar
'''

s3_conn = S3('s3_conn')
base_boto = s3_conn.get_conn()


def execute_commands_on_linux_instances(**kwargs):
    """Runs commands on remote linux instances
    :param connection: a boto/boto3 ssm client
    :param command: a list of strings, each one a command to execute on the instances
    :param id_select: a list of instance_id strings, of the instances on which to execute the command
    :return: the response from the send_command function (check the boto3 docs for ssm client.send_command() )
    """

    client = kwargs['connection']

    resp = client.send_command(
        DocumentName="AWS-RunShellScript",  # One of AWS' preconfigured documents
        Parameters={'commands': kwargs['command']},
        InstanceIds=kwargs['id_select'],
    )
    return resp

# Example use:


ftp_s3_sync = PythonOperator(
    task_id='ftp_s3',
    python_callable=execute_commands_on_linux_instances,
    provide_context=True,
    op_kwargs={
        'command': bash,
        'connection': base_boto,
        'id_select': 'connection ids go here'
    }
)
ssm_client = boto3.client('ssm')  # Need your credentials here
commands = ['echo "hello world"']
instance_ids = ['an_instance_id_string']
execute_commands_on_linux_instances(ssm_client, commands, instance_ids)
