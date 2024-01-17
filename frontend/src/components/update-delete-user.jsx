import { useEffect, useState } from "react";
import { deleteUser, getUsers, updateUser } from "../api/user";
import Table from "./table";
import Form from "./form";
import Toast from "./toast";

const Modal = ({ isOpen, onClose, user, updateUser, handleFieldChange }) => {
  const formFields = [
    { name: "name", label: "Name", default: user?.name },
    { name: "email", label: "Email", default: user?.email },
    { name: "age", label: "Age", type: "number", default: user?.age },
    {
      name: "birthday",
      label: "Birthday",
      type: "date",
      default: user?.birthday,
    },
    {
      name: "member_number",
      label: "Member Number",
      default: user?.member_number,
    },
    {
      name: "activated_at",
      label: "Activated At",
      type: "date",
      default: user?.activated_at,
    },
  ];
  return (
    <>
      {isOpen && (
        <div className="fixed  inset-0 z-50 flex items-center justify-center">
          <div className="absolute inset-0 bg-gray-800 opacity-75"></div>
          <div className="bg-white w-[60%] p-8 rounded shadow-md z-50">
            <div className="flex flex-row justify-end">
              <button
                onClick={onClose}
                className="text-gray-500 hover:text-gray-700"
              >
                X
              </button>
            </div>
            <Form
              fields={formFields}
              onSubmit={updateUser}
              onFieldChange={handleFieldChange}
            />
          </div>
        </div>
      )}
    </>
  );
};

export default function UpdateDeleteUser() {
  const [page, setPage] = useState(1);
  const [limit, setLimit] = useState(10);
  const [headers, setHeaders] = useState([]);
  const [data, setData] = useState([]);

  const [userData, setUserData] = useState({
    name: "",
    email: "",
    age: 0,
    birthday: "",
    member_number: "",
    activated_at: "",
  });

  const handleFieldChange = ({ name, value }) => {
    setUserData((prevData) => ({
      ...prevData,
      [name]: value,
    }));
  };

  const [isModalOpen, setModalOpen] = useState(false);

  const update = async () => {
    try {
      const request_start_at = performance.now();
      // Handle form submission logic here
      const response = await updateUser(userData, userData.id);
      const request_end_at = performance.now();
      const request_duration = request_end_at - request_start_at;

      Toast({
        message: `${response?.data?.message} - ${
          request_duration.toFixed(1) * 100
        }ms`,
        type: "success",
      });
      closeModal();
      getAllUsers(page, limit);
    } catch (error) {
      Toast({
        message:
          error?.response?.data?.error?.toUpperCase() || "Something went wrong",
        type: "error",
      });
    }
  };

  const deleteUsr = async (id) => {
    try {
      const request_start_at = performance.now();
      // Handle form submission logic here
      const response = await deleteUser(id);
      const request_end_at = performance.now();
      const request_duration = request_end_at - request_start_at;

      Toast({
        message: `${response?.data?.message} - ${
          request_duration.toFixed(1) * 100
        }ms`,
        type: "success",
      });
      getAllUsers(page, limit);
    } catch (error) {
      Toast({
        message:
          error?.response?.data?.error?.toUpperCase() || "Something went wrong",
        type: "error",
      });
    }
  };

  const openModal = (user) => {
    setUserData(user);
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
  };

  useEffect(() => {
    getAllUsers(page, limit);
  }, [page, limit, userData]);

  async function getAllUsers(page, limit) {
    try {
      const response = await getUsers(page, limit);
      //   setUsers(response.data.data);
      setHeaders(Object.keys(response?.data?.data?.users[0]));
      setData(response?.data?.data?.users);
    } catch (error) {
      Toast({
        message: error?.response?.data?.error?.toUpperCase(),
        type: "error",
      });
    }
  }
  const previous = () => {
    setPage((value) => (value > 1 ? value - 1 : 1));
  };

  const next = () => {
    alert(data.length);
    setPage((value) => (data.length == 0 ? value : value + 1));
  };
  return (
    <div className="flex flex-col mt-8 gap-4">
      <Toast />
      <div className="w-full">
        <div className="flex flex-row justify-between">
          <div>
            <label htmlFor="limit">No of Records</label>
            <select
              value={limit}
              onChange={(event) => {
                setLimit(event.target.value);
              }}
              className="border rounded ml-4"
              id="limit"
            >
              <option>10</option>
              <option>25</option>
              <option>50</option>
              <option>100</option>
            </select>
          </div>
          <div className="flex flex-row gap-4">
            <button onClick={previous}>Previous</button>
            <button onClick={next}>Next</button>
          </div>
        </div>
        <div className="mb-2"></div>
        <Table
          headers={headers}
          data={data}
          updateUser={openModal}
          deleteUser={deleteUsr}
        />
      </div>

      <Modal
        isOpen={isModalOpen}
        onClose={closeModal}
        user={userData}
        updateUser={update}
        handleFieldChange={handleFieldChange}
      ></Modal>
    </div>
  );
}
