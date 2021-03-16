async function readHistory(isbns) {

    // Get the historical status data for the ISBNs.
    let history;
    let promise = swaggerClient
        .then(
            client => client.apis.api.bookHistory({isbns: isbns}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookHistoryResult => history = JSON.parse(bookHistoryResult.data),
            reason => console.error('failed api call: ' + reason)
        );
    await promise;

    return history;
}

async function buildHistory(history) {

    // Remove all historical status data.
    removeAllChildNodes(historyModalList);

    // Iterate through the historical status data.
    let first = true;
    history.history.reverse().forEach(function (status) {

        // Create the status entry.
        let listItem = document.createElement("li");
        listItem.classList.add("list-group-item");
        if (first) {
            first = false;
            listItem.classList.add("active");
            listItem.setAttribute("aria-current", "true");
        }

        // Create a div for the Status data and time entry as columns.
        let statusDiv = document.createElement("div");
        statusDiv.classList.add("col");
        statusDiv.innerHTML = status.type;
        let timeDiv = document.createElement("div");
        timeDiv.classList.add("col");
        timeDiv.innerHTML = status.time;

        // Create the row the columns will live in.
        let row = document.createElement("div");
        row.classList.add("row");
        row.classList.add("justify-content-between");
        row.appendChild(statusDiv);
        row.appendChild(timeDiv);

        // Add the row to the list item.
        listItem.appendChild(row);

        // Append the status to the parent.
        historyModalList.appendChild(listItem)
    });
}
