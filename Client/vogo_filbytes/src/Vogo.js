import React from 'react';

const Header = ({Name}) => {
  return (
    <div style={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center', padding: '10px' }}>
      <h1 style={{ color: 'maroon' }}>{Name}</h1>
      <div style={{ marginLeft: 'auto' }}>ID</div>
    </div>
  );
};
export default Header;