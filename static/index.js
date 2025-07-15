const form = document.getElementById("regression-form");
form.addEventListener('submit', getFormData);


//This function will pull data from the html form, parse it to an object and send (POST) it to our Go Server
function getFormData(event) {
    event.preventDefault();
    const skillData = document.getElementById("skill").value;
    console.log(skillData)
    const lastPracticedData = document.getElementById("lastpracticed").value;
    console.log(lastPracticedData)
    const complexity = document.getElementById("complexity").value;
    console.log(complexity)

    const formObject = { skill: skillData, lastpracticed: lastPracticedData, complexity: complexity};

    const myJSON = JSON.stringify(formObject)

    //Sending the Data to the backend and then receiving the json
    fetch('/api', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: myJSON
    })
    .then(response => response.json())
    .then(responseStruct => {
    console.log(responseStruct.regressorscore);
    console.log(responseStruct.interpretation);
    });

    

}

