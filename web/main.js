var myRequest = new XMLHttpRequest();

myRequest.responseType = "json";
myRequest.open("GET", "http://localhost:8080/people", true);
myRequest.send();
