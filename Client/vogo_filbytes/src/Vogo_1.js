import React from 'react';

const FourColumnsWithProps = ({ column1, column2, column3, column4 }) => {
    const boxStyle = {
        flex: 1,
        border: '4px solid maroon',
        padding: '10px',
        fontWeight: 'bold',
        backgroundColor:'#f2f2f2',
      };
      const containerStyle = {
        display: 'flex',
        justifyContent: 'space-between',
        padding: '20px',
        width: '50%', // Set a fixed width for the container
        height: '100%', // Set a fixed height for the container
        margin: 'auto',
      };
    
  return (
    <div style={containerStyle}>
      <div style={boxStyle}>{column1}</div>
      <div style={boxStyle}>{column2}</div>
      <div style={boxStyle}>{column3}</div>
      <div style={boxStyle}>{column4}</div>
    </div>
  );
};

export default FourColumnsWithProps;
