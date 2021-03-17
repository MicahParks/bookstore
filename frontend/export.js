function download(filename, text) {
    let element = document.createElement('a');
    element.setAttribute('href', 'data:text/plain;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    element.style.display = 'none';
    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}

async function exportCSV() {

    // Get the CSV export.
    let csv;
    let promise = swaggerClient
        .then(
            client => client.apis.api.bookCSV(),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookReadResult => csv = bookReadResult.data,
            reason => showAlert(reason)
        );
    await promise;

    download("export.csv", csv);
}

async function exportJSON() {
    readHistory().then(function (history) {
        readBooks().then(function (books) {
            download("export.json", JSON.stringify({
                history: history,
                books: books
            }));
        });
    });
}
