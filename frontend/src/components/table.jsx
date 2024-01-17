const Table = ({ headers, data, updateUser, deleteUser }) => {
  return (
    <table className="w-full responsive bg-white border border-gray-300">
      <thead>
        <tr>
          {/* {JSON.stringify(headers)} */}
          {headers.map((header) => (
            <th key={header} className="py-2 px-4 border-b uppercase">
              {header.replace("_", " ")}
            </th>
          ))}
          <th className="py-2 px-4 border-b">ACTIONS</th>

          {/* <th className="py-2 px-4 border-b">ID</th>
          <th className="py-2 px-4 border-b">Name</th>
          <th className="py-2 px-4 border-b">Age</th> */}
        </tr>
      </thead>
      <tbody>
        {data?.map((row) => (
          <tr key={row.id}>
            {headers.map((header) => (
              <td key={header} className="py-2 px-4 border-b">
                {row[header]}
              </td>
            ))}
            <td className=" py-2 px-4 border-b">
              <button onClick={() => updateUser(row)} className="bg-yellow-500 hover:bg-yellow-700 text-white font-bold py-2 px-4 rounded w-full my-2">
                Update
              </button>
              <button onClick={() => deleteUser(row.id)} className="bg-red-500 hover:bg-red-700 text-white font-bold py-2 px-4 rounded w-full my-2">
                Delete
              </button>
            </td>
          </tr>
        ))}
      </tbody>
    </table>
  );
};

export default Table;
