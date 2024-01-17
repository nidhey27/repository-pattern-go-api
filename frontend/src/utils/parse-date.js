const dateRegex = /^\d{4}-\d{2}-\d{2}$/;
function convertDate(inputDate) {
    // Parse the input date string
    const parsedDate = new Date(inputDate);

    // Format the date as per the desired output format
    const outputDate = parsedDate.toISOString().split('T')[0] + 'T00:00:00Z';

    return outputDate;
}

export { convertDate, dateRegex }