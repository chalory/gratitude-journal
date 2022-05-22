import React from 'react';
import '../App.css';

const Notes = ({ data }) => {
  return (
    <div className='col-xs-4'>
        <ul>
            {data && data.map((item, index) => <li key={index}> <h3>{item.user_name}'s post</h3> <p>{item.date}</p> {item.entry}</li>)}
        </ul>
    </div>
  );
};

export default Notes;