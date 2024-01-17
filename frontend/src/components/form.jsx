import { useState } from "react";
import { dateRegex, convertDate } from "../utils/parse-date";
const Form = ({ fields, onSubmit, onFieldChange }) => {
  const [formData, setFormData] = useState(
    fields.reduce((acc, field) => {
      acc[field.name] = field?.default || "";
      return acc;
    }, {})
  );

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({
      ...prevData,
      [name]:  dateRegex.test(value) ? convertDate(value) : !isNaN(value) ? Number(value) : value,
    }));
    // Send updated data to parent component
    onFieldChange({ name, value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
  };

  return (
    <div className="max-w-md mx-auto p-8 border rounded-md">
      <h2 className="text-2xl font-bold mb-4">User Information</h2>
      <form onSubmit={handleSubmit}>
        {fields.map((field) => (
          <div key={field.name} className="mb-4">
            <label className="block text-sm font-medium text-gray-600">
              {field.label}
            </label>
            <input
              type={field.type || "text"}
              name={field.name}
              value={formData[field.name]}
              onChange={handleChange}
              className="mt-1 p-2 border rounded-md w-full"
            />
          </div>
        ))}

        <div className="mt-4">
          <button
            type="submit"
            className="bg-blue-500 text-white py-2 px-4 rounded-md hover:bg-blue-600"
          >
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default Form;
