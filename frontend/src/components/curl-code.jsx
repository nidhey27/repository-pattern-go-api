const CurlCommand = ({ method, apiURL, requestBody }) => {

  const generateCurlCommand = () => {
    switch (method) {
      case "GET":
        return `curl --location --request GET '${apiURL}'`;
      case "POST":
        return `curl --location '${apiURL}' \\
                --header 'Content-Type: application/json' \\
                --data-raw '${requestBody}'`;
      case "PUT":
        return `curl --location --request PUT '${apiURL}' \\
                --header 'Content-Type: application/json' \\
                --request PUT \\
                --data-raw '${requestBody}'`;
      case "DELETE":
        return `curl --location --request DELETE '${apiURL}' \\
                --request DELETE`;
      default:
        return "";
    }
  };

  const curlCommand = generateCurlCommand();

  return (
    <div>
      <label htmlFor="curl-request">cURL Command:</label>
      <div id="curl-request" className="overflow-scroll">
        <code>{curlCommand}</code>
      </div>
    </div>
  );
};

export default CurlCommand;
