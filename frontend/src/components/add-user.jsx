import { useState } from "react";
import Form from "./form";
import { createUser } from "../api/user";
import CurlCommand from "./curl-code";
import JsonFormatter from "react-json-formatter";
import Toast from "./toast";

export default function AddUser() {
  const [response, setResponse] = useState({});
  const [formData, setFormData] = useState({
    name: "",
    email: "",
    age: "",
    birthday: "",
    member_number: "",
    activated_at: "",
  });
  const handleSubmit = async (formData) => {
    try {
      const request_start_at = performance.now();
      // Handle form submission logic here
      let response = await createUser(formData);
      const request_end_at = performance.now();
      const request_duration = request_end_at - request_start_at;
      setResponse(response);
      Toast({
        message: `User Created - ${request_duration.toFixed(1) * 100}ms`,
        type: "success",
      });
    } catch (error) {
      setResponse(error.response.data);
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
    { name: "name", label: "Name" },
    { name: "email", label: "Email" },
    { name: "age", label: "Age", type: "number" },
    { name: "birthday", label: "Birthday", type: "date" },
    { name: "member_number", label: "Member Number" },
    { name: "activated_at", label: "Activated At", type: "date" },
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
        <label htmlFor="curl-request"></label>
        <CurlCommand
          method="POST"
          apiURL={import.meta.env.VITE_BACKEND_URL + `/user`}
          requestBody={JSON.stringify(formData)}
        />

        <div className="border-t py-4">
          <h1 className="font-[500]">Response</h1>
          <div className="response">
            <div className="flex flex-col">
              <h1>
                HTTP Status: {response?.code || response?.response?.data?.code}
              </h1>
              <h1>
                Response Size: {response?.headers?.["content-length"] || 0} B
              </h1>
            </div>
            <JsonFormatter
              json={response?.data?.data || response}
              tabWith={4}
              jsonStyle={jsonStyle}
            />
          </div>
        </div>
      </div>
    </div>
  );
}
