## Tech Challenge #002

<img src="divido-tech-challenge-002.jpg" width="50%" alt="Divido Tech Challenge #002">

We are going to write a function that takes a list of dishes that our new chef can cook and builds a menu for our customers.<br/>

You can use **_any programming language_** of your choice to complete this challenge and for this reason your function will 
be judged & tested manually once you submit a pull request (there won't be automated tests).<br/>

There is a `build-menu.go` file and a `build-menu_test.go` file for you to look through, but you **don't** have to code the solution in Go.<br/>
You can create your own file and code your solution any way you like as long as you maintain the equivalent function signature in your given language of choice.<br/>

For example, in Java you would have a file called `BuildMenu.java` that would look like this:

```
public class Solution {
    public String buildMenu(Map<String, List<String>> courseItems, List<String> chefDishes, String restaurantName) {
        // Your solution goes here
        return "";
    }
}
```

The function will be passed the 3 following parameters:

1. A map of the various courses
```
{
    "antipasti": {
        "tacos", "chips", "plantain", "quesadilla",
    },

    "mains": {
        "steak", "nachos", "chicken", "pizza",
        "sandwich", "taco", "sushi", "burger",
        "pork", "pasta", "wrap", "rice",
    },

    "desserts": {
        "pears", "churros", "cheese",
        "icecream", "sorbet", "cupcake",
    },
}
```

2. A list of the dishes our chef is capable of cooking
```
{ "churros", "tacos", "pasta", "rice", "pears", "plantain", "chips", "cheese", "nachos", "sushi" }
```

3. The name of the restaurant<br/>
```Nacho Average Mexican Diner```

Your solution must return the following output string:
```
**** Nacho Average Mexican Diner ****
             - Menu -
Antipasti
tacos, chips or plantain
Mains
nachos, sushi, pasta or rice
Desserts
pears, churros or cheese
```

### Submitting Your Solution
To submit your solution to the challenge, create a branch with your GitHub username,
followed by a _forward slash_ and the name of the challenge (e.g. `mygithubusername/build-menu`).<br/>
Then submit a pull request.

**Tip:** Be careful of formatting, your solution must match the expected output down to the last character.