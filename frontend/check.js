async function checkin(button) {

    // Get the ISBN from the button.
    let isbns = [button.getAttribute('data-bs-isbn')];

    // Check out the given ISBNs.
    return swaggerClient
        .then(
            client => client.apis.api.bookCheckin({isbns: isbns}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookCheckinResult => { /* No operation.*/
            },
            reason => console.error('failed api call: ' + reason)
        );
}

async function checkout(button) {

    // Get the ISBN from the button.
    let isbns = [button.getAttribute('data-bs-isbn')];

    // Check out the given ISBNs.
    return swaggerClient
        .then(
            client => client.apis.api.bookCheckout({isbns: isbns}),
            reason => console.error('failed to load the spec: ' + reason)
        )
        .then(
            bookCheckoutResult => { /* No operation.*/
            },
            reason => console.error('failed api call: ' + reason)
        );
}
