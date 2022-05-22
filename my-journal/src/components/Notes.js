import React from 'react';

const Notes = ({ data }) => {
  return (
    <div>
      <ul>
        {data && data.map((item, index) => <li key={index}>{item.id}, {item.entry}, {item.date}, {item.user_name}</li>)}
      </ul>
    </div>
  );
};

export default Notes;