const form = document.getElementById("tojs");
form.addEventListener('submit', function(e) {
    e.preventDefault();
    let url = "http://localhost:8080/test/new-url-string";
    let data = document.getElementById("urltwo").value;
    fetch(url, {
        method: 'POST',
        body: JSON.stringify(data),
        headers: {
            "Content-Type": "application/json",
        }

    })
    .then(function(response) {
        return response.json()
    })
    .then(function(data) {
        console.log(data)
        document.getElementById("showShort").innerHTML = "http://localhost:8080/" + data;
    })
});
