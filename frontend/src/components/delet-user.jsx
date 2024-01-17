import { useState } from "react";
import Form from "./form";
import { getUserByID } from "../api/user";
import JsonFormatter from "react-json-formatter";
import CurlCommand from "./curl-code";
export default function DeleteUser() {
  const [response, setResponse] = useState({});
  const [formData, setFormData] = useState({
    id: "",
  });
  const handleSubmit = async (formData) => {
    const request_start_at = performance.now();

    let response = await getUserByID(formData.id);

    const request_end_at = performance.now();
    const request_duration = request_end_at - request_start_at;
    setResponse(response);
    alert(request_duration + "ms");
  };
  const handleFieldChange = ({ name, value }) => {
    setFormData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };
  const formFields = [{ name: "id", label: "ID" }];

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
      <div className="w-1/2">
        <Form
          fields={formFields}
          onSubmit={handleSubmit}
          onFieldChange={handleFieldChange}
        />
      </div>
      <div className="w-1/2 flex flex-col gap-4">
        <CurlCommand
          method="DELETE"
          apiURL={import.meta.env.VITE_BACKEND_URL + `/user/${formData.id}`}
        />

        <div className="border-t py-4">
          <h1 className="font-[500]">Response</h1>
          {response != null && (
            <div className="response">
              <div className="flex flex-col">
                {/* <h1>HTTP Status: {response.data.code || response.response.data.code}</h1>
                {/* <h1>Response Size: {response.headers?.["content-length"]} B</h1> */}
              </div>
              <JsonFormatter
                json={response.data.data}
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
