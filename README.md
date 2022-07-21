# coffee-machine
Requirement:
Design a scalable coffee machine with the following functionalities. 
1. The machine at the start should display the list of raw materials (coffee, water, milk, chocolate powder, etc) it needs. 
2. The user should able to input the raw materials and their respective quantities.
3. Should display the coffee options available with it - eg mocha, cappuccino, etc. 
4. It should also have an option to display the current quantity of raw material available.
5. should mark an item as unavailable if any of the ingredients have finished (or are less than required quantity).
6. Should display the quantities of items available when the given option is selected
 
It should have apis for the above functionalities and should display appropriate messages while performing any activity.



Design:

coffee-machine shouls have the following API's
1. coffee-machine/startMachine
Request : "/GET"


Response :
{
    "<rawmaterial-1>":"quantity",
    "<rawmaterial-2>":"quantity",
    ...
}

2. coffee-machine/addRawMaterial
Request : "/POST"
{
    "<rawmaterial-1>":"quantity",
    "<rawmaterial-2>":"quantity",
    ...
}
Response : 
If addes :
{
    "status":"SUCCESS",
    "message":"Raw materials added"
}

If not available :
{
    "status":"FAILED",
    "message":"Unable to add raw materials.Please try again after restarting the machine."
}


3. coffee-machine/coffeeTypes
Request : "/GET"

Response :
{
    "<coffeetype-1>":"quantity",
    "<coffeetype-2>":"quantity",
    ...
}

4. coffee-machine/rawMaterialsAvailable
Request : "/GET"

Response :
{
    "<rawmaterial-1>":"quantity",
    "<rawmaterial-2>":"quantity",
    ...
}

5. coffee-machine/makeCoffee
Request : "/POST"
{
    "coffeeType" : "<Selected coffee type by the user>"
}
Response : 
If available :
{
    "status":"SUCCESS",
    "message":"Availabe"
}

If not available :
{
    "status":"FAILED",
    "message":"Currently <User inputted coffee type> is unavailabe. Please select different type of coffee"
}


6. coffee-machine/coffeeTypesAvailable
Request : "/GET"

Response :
{
    "<coffeetype-1>":"quantity",
    "<coffeetype-2>":"quantity",
    ...
}