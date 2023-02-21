import React, { useState } from 'react';

const Tab = (props) => {
  const { activeTab, label, onClick } = props;
  let className = 'tab-list-item';

  if (activeTab === label) {
    className += ' tab-list-active';
  }

  return (
    <li className={className} onClick={() => onClick(label)}>
      {label}
    </li>
  );
};

const TabContent = (props) => {
  const { activeTab } = props;

  const tabContentMap = {
    Home: <div>Home Tab Content</div>,
    About: <div>About Tab Content</div>,
    Filbytes: <div>Filbytes Tab Content</div>,
    Login: <div>Login Tab Content</div>,
  };

  return <div>{tabContentMap[activeTab]}</div>;
};

const Tabs = () => {
  const [activeTab, setActiveTab] = useState('Home');

  const handleClick = (label) => {
    setActiveTab(label);
  };

  return (
    <div>
      <ul className="tab-list">
        <Tab activeTab={activeTab} label="Home" onClick={handleClick} />
        <Tab activeTab={activeTab} label="About" onClick={handleClick} />
        <Tab activeTab={activeTab} label="Filbytes" onClick={handleClick} />
        <Tab activeTab={activeTab} label="Login" onClick={handleClick} />
      </ul>
      <div className="tab-content">
        <TabContent activeTab={activeTab} />
      </div>
    </div>
  );
};

export default Tabs;
