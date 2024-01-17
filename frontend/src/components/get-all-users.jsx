import { useState } from "react";
import Form from "./form";
import { getUsers } from "../api/user";
import JsonFormatter from "react-json-formatter";
import CurlCommand from "./curl-code";
import Toast from "./toast";
export default function GetUsers() {
  const [response, setResponse] = useState({});
  const [formData, setFormData] = useState({
    page: 1,
    limit: 10,
  });
  const handleSubmit = async (formData) => {
    try {
      const request_start_at = performance.now();
      // Handle form submission logic here
      let response = await getUsers(formData.page, formData.limit);
      const request_end_at = performance.now();
      const request_duration = request_end_at - request_start_at;
      setResponse(response);
      Toast({
        message: `Users Fetched - ${request_duration.toFixed(1) * 100}ms`,
        type: "success",
      });
    } catch (error) {
      Toast({
        message: error.response.data.error.toUpperCase(),
        type: "error",
      });
    }
  };
  const handleFieldChange = ({ name, value }) => {
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };
  const formFields = [
    { name: "page", label: "Page", default: 1 },
    { name: "limit", label: "Limit", default: 10 },
  ];

  const jsonStyle = {
    propertyStyle: { color: "cyan" },
    stringStyle: { color: "lightgreen" },
    numberStyle: { color: "darkorange" },
    braceStyle: {
      color: "white",
    },
  };

  return (
    <div className="flex flex-row mt-8">
      <Toast />
      <div className="w-1/2">
        <Form
          fields={formFields}
          onSubmit={handleSubmit}
          onFieldChange={handleFieldChange}
        />
      </div>
      <div className="w-1/2 flex flex-col gap-4">
        <CurlCommand
          method="GET"
          apiURL={
            import.meta.env.VITE_BACKEND_URL +
            `/user?page=${formData.page}&limit=${formData.limit}`
          }
        />

        <div className="border-t py-4">
          <h1 className="font-[500]">Response</h1>
          {response != null && (
            <div className="response">
              <div className="flex flex-col">
                <h1>HTTP Status: {response?.data?.code}</h1>
                <h1>
                  Response Size: {response?.headers?.["content-length"]} B
                </h1>
              </div>
              <JsonFormatter
                json={response?.data?.data || response}
                tabWith={4}
                jsonStyle={jsonStyle}
              />
            </div>
          )}
        </div>
      </div>
    </div>
  );
}
