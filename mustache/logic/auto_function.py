# %% testing basic mustache
import json
import chevron


with open('templates/hello.mustache', 'r') as f:
    world = chevron.render(f, {'mustache': 'World'})
    print(world)

# %% mustache list
import chevron
input_template = 'templates/functions.mustache'
output_file = 'out/basic_function_mustache.py'

function_names = {
    "functions": [{
        "name": "mustache_one",
        "message": "Hello from #1"},
        {
        "name": "mustache_two",
        "message": "Hello from #2"},
        {
        "name": "mustache_three",
        "message": "Hello from #3"}
    ]
}
# load in a yaml file and render it to a python object


with open(input_template, 'r') as file_read:
    functions = chevron.render(file_read, function_names)
    with open(output_file, 'w') as file_write:
        file_write.write(functions)

# Yes following is bad. Just for demo.
exec(open(output_file).read())

# Notes: Requires use of outside library and learning of syntax but provides
# slightly better dev experience once learned


# %% python file builder
output_file = 'out/basic_function_python.py'


def generate_function(function_name: str, message: str) -> str:
    function_string = f'''
def {function_name}():
    print("My function Name is: {function_name}")
    print("My Message is: {message}")
    print('')
{function_name}()
    '''

    return function_string


function_names = {
    "functions": [{
        "name": "python_one",
        "message": "Hello from #1"},
        {
        "name": "python_two",
        "message": "Hello from #2"},
        {
        "name": "python_three",
        "message": "Hello from #3"}
    ]
}

with open(output_file, 'w') as f:
    for functions in function_names['functions']:
        f.write(generate_function(
            function_name=functions['name'],
            message=functions['message']))

# Yes following is bad. Just for demo.
exec(open(output_file).read())

# Notes: Leverages python syntax which is more familiar and doesnt require
# additional tooling. Would consider that string formatting is an issue with
# python but that exists in mustache as well so no benefit either way.

# For both its pretty easy to write out the python function you need or look
# at an existing one and write out the template.
# %% Use openapi 3.0 yaml/json to generate partial of service stack
data_file = "data/petstore.json"
with open(data_file, 'r') as file:
    data = json.load(file)


cleaned_dict = {}
cleaned_dict['ServiceName'] = data['info']['title'].replace(' ', '')


mustache_template = 'templates/mockservice.mustache'
output_file = 'out/mockservice_mustache.py'
with open(mustache_template, 'r') as file_read:
    filled_form = chevron.render(file_read,  cleaned_dict)
    with open(output_file, 'w') as file_write:
        file_write.write(filled_form)


# %%


 The initial goals are to reduce the amount of boilerplate code that needs to be hand written, allow for fluid changes to a services api structure and overall speed up the time from api design to service deploy. 