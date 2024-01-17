// Tabs.jsx

import { useState } from "react";

const Tabs = ({ tabs }) => {
  const [activeTab, setActiveTab] = useState(tabs[0]?.id);

  const changeTab = (tabId) => {
    setActiveTab(tabId);
  };

  return (
    <div className="w-full ">
      <div className="flex ">
        {tabs.map((tab) => (
          <button
            key={tab.id}
            className={`px-4 py-2 ${
              activeTab === tab.id
                ? "bg-blue-500 text-white rounded-xl mx-1"
                : "bg-gray-300 text-gray-700 rounded-xl mx-1"
            }`}
            onClick={() => changeTab(tab.id)}
          >
            {tab.label}
          </button>
        ))}
      </div>
      <div className="mt-4">
        {tabs.map((tab) => (
          <div
            key={tab.id}
            className={`${activeTab === tab.id ? "" : "hidden"}`}
          >
            <tab.component />
          </div>
        ))}
      </div>
    </div>
  );
};

export default Tabs;
