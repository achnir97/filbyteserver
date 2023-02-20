import React from 'react';

const MyGrid = () => {
  const containerStyle = {
    display: 'flex',
  };
  
  const leftColumnStyle = {
    flex: '0 0 50%',
    backgroundColor: '#f2f2f2',
    padding: '20px',
  };
  
  const rightColumnStyle = {
    flex: '0 0 50%',
    backgroundColor: '#ddd',
    padding: '20px',
    display: 'flex',
    flexDirection: 'column',
  };
  
  const topRightRowStyle = {
    flex: '1',
    backgroundColor: '#eee',
    marginBottom: '10px',
  };
  
  const bottomRightRowStyle = {
    flex: '1',
    backgroundColor: '#ccc',
  };

  return (
    <div style={containerStyle}>
      <div style={leftColumnStyle}>
        <p>This is the left column</p>
      </div>
      <div style={rightColumnStyle}>
        <div style={topRightRowStyle}>
          <p>This is the top right row</p>
        </div>
        <div style={bottomRightRowStyle}>
          <p>This is the bottom right row</p>
        </div>
      </div>
    </div>
  );
};

export default MyGrid;
