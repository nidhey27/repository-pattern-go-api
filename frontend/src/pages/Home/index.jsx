import AddUser from "../../components/add-user";
import GetUsers from "../../components/get-all-users";
import GetUser from "../../components/get-user";
import Tabs from "../../components/tabs";
import UpdateDeleteUser from "../../components/update-delete-user";

export default function HomePage() {
  const tabs = [
    { label: "Add User", id: "Tab 1", component: AddUser },
    { label: "Get user by ID", id: "Tab 2", component: GetUser },
    { label: "Get All users", id: "Tab 3", component: GetUsers },
    { label: "Update & Delete User", id: "Tab 4", component: UpdateDeleteUser },
  ];

  return (
    <main className="max-w-7xl mx-auto">
      <div className="shadow-lg p-6 rounded-2xl">
        <h1 className="text-xl font-bold">User Management</h1>
        <div className="mb-5"></div>
        <Tabs tabs={tabs} />
      </div>
    </main>
  );
}
