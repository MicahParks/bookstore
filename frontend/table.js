async function refresh(isbns) {

    let tableData = {
        books: null,
        statuses: null
    };

    // Get the Book data for the ISBNs.
    let promise = readBooks(isbns).then(function (books) {
        tableData.books = books;
    });
    await promise;

    // Get the Status data for the ISBNs.
    promise = swaggerClient
        .then(
            client => client.apis.api.bookStatus({isbns: isbns}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookStatusResult => tableData.statuses = JSON.parse(bookStatusResult.data),
            reason => showAlert(reason)
        );
    await promise;

    return tableData;
}

async function buildTable() {

    // Remove all rows in the table.
    removeAllChildNodes(table);

    // Get the latest status data.
    return refresh(null).then(function (tableData) {

        // Iterate through all books.
        let index = 0;
        for (const [isbn, book] of Object.entries(tableData.books)) {

            // Create a new row for the book based off of the template.
            let row = rowTemplate.cloneNode(true);

            // Assign the row's ID to the ISBN.
            row.id = 'row' + index;
            index++;

            // Get the Status data for the ISBN.
            let status = tableData.statuses[isbn]

            // Decide what the status button should be.
            let available = status.available;
            let unavailable = status.unavailable;
            if (available === undefined) {
                available = 0;
            }
            if (unavailable === undefined) {
                unavailable = 0;
            }
            let availableButton = document.createElement("button");
            availableButton.type = "button";
            availableButton.disabled = true;
            availableButton.classList.add("btn");
            availableButton.classList.add("btn-success");
            availableButton.innerHTML = available;
            let unavailableButton = document.createElement("button");
            unavailableButton.type = "button";
            unavailableButton.disabled = true;
            unavailableButton.classList.add("btn");
            unavailableButton.classList.add("btn-danger");
            unavailableButton.innerHTML = unavailable;

            // Assign the columns for the row.
            row.cells[0].innerHTML = isbn;
            row.cells[1].appendChild(availableButton);
            row.cells[1].appendChild(unavailableButton);
            row.cells[2].innerHTML = book.title;
            row.cells[3].innerHTML = book.author;
            row.cells[4].innerHTML = book.description;

            // Add the row to the table.
            table.appendChild(row);

            // Label the row's buttons with the its ISBN.
            for (let button of $("#" + row.id + " :button")) {
                button.setAttribute("data-bs-isbn", isbn);
            }
        }
    })
}
