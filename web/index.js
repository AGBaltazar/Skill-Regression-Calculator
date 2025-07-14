

//This function will pull data from the html form, parse it to an object and send (POST) it to our Go Server
function getFormData() {
    const skillData = document.getElementById("skill");
    const lastPracticedData = document.getElementById("laspracticed");
    const complexity = document.getElementById("complexity");

    const formObject = { skill: skillData.value, lastpracticed: lastPracticedData.value, complexity: complexity.value};

    const myJSON = JSON.stringify(formObject)

    fetch('/api/data', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: myJSON
})

    .then(response => response.json())//parsing response from Go
    .then(console.log("Post sucessfully sent!"))

}