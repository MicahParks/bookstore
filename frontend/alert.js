// https://www.w3docs.com/snippets/javascript/how-to-create-a-new-dom-element-from-html-string.html

function showAlert(text) {

    // The alert HTML.
    let alert = '<div class="alert alert-danger alert-dismissible fade show" role="alert"><button type="button" class="btn-close" data-bs-dismiss="alert" aria-label="Close"></button></div>';

    // Turn the alert into a div element.
    let temp = document.createElement("template");
    temp.innerHTML = alert;
    let alertDiv = temp.content.firstChild;

    // Add the alert text to the div.
    alertDiv.appendChild(document.createTextNode(String(text.response.body.message).trim()));

    // Remove any previous alert and replace it with this one.
    let div = document.getElementById("alertDiv");
    removeAllChildNodes(div);
    div.appendChild(alertDiv);
}
