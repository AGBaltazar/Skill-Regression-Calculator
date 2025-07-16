const form = document.getElementById("regression-form");
form.addEventListener('submit', getFormData);

// This function will pull data from the html form, parse it to an object and send (POST) it to our Go Server
function getFormData(event) {
    event.preventDefault();

    const skillData = document.getElementById("skill").value;
    const lastPracticedData = document.getElementById("lastpracticed").value;
    const complexity = document.getElementById("complexity").value;
    const formObject = {
        skill: skillData,
        lastpracticed: lastPracticedData,
        complexity: complexity
    };

    const myJSON = JSON.stringify(formObject);

    // Sending the Data to the backend and then receiving the json
    fetch('/api', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: myJSON
    })
    .then(response => response.json())
    .then(responseStruct => {
        document.getElementById('retention-score').innerHTML = `<p>${responseStruct.regressorscore}</p>`;
        document.getElementById('interpretation').innerHTML = `<p>${responseStruct.interpretation}</p>`;
        const videoContainer = document.getElementById('video-links');
        videoContainer.innerHTML = "";

        if (Array.isArray(responseStruct.youtubelinks)) {
            responseStruct.youtubelinks.forEach(link => {
                const a = document.createElement('a');
                a.href = link;
                a.textContent = link;
                a.target = "_blank";
                a.className = "btn btn-outline-primary btn-sm";
                videoContainer.appendChild(a);
            });
        } else {
            console.warn("No YouTube links received");
        }
    })
    .catch(error => {
        console.error("Error fetching data:", error);
    });
}
