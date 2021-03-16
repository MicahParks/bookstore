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
            reason => console.error('failed api call: ' + reason)
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
        for (const [isbn, book] of Object.entries(tableData.books)) {

            // Create a new row for the book based off of the template.
            let row = rowTemplate.cloneNode(true);

            // Assign the row's ID to the ISBN.
            row.id = isbn;

            // Get the Status data for the ISBN.
            let status = tableData.statuses[isbn]

            // Decide what the status button should be.
            let button = document.createElement("button");
            button.type = "button";
            button.disabled = true;
            switch (status.type) {
                case "acquired":
                    button.classList.add("btn");
                    button.classList.add("btn-success");
                    button.innerHTML = "Acquired";
                    break;
                case "checkin":
                    button.classList.add("btn");
                    button.classList.add("btn-primary");
                    button.innerHTML = "Checked in";
                    break;
                case "checkout":
                    button.classList.add("btn");
                    button.classList.add("btn-secondary");
                    button.innerHTML = "Checked out";
                    break;
            }

            // Assign the columns for the row.
            row.cells[0].innerHTML = isbn;
            row.cells[1].appendChild(button);
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
