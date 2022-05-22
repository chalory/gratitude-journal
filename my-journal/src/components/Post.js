import React, { props, useRef } from 'react';
import axios from 'axios';

const Post = () => {
    const id = useRef(0);
    const entry = useRef(null);
    const date = useRef(0);
    const user_name = useRef(null);

    // const getInput = () => {
    //     return (
    //         <div>
    //         <form onSubmit={props.onSubmit}>
    //           <input type="text" value={props.userName} onChange={props.changed} />
    //           <input type="submit" value="Submit" />
    //         </form>
    //       </div>
    //     )
    // };

    axios.post('/entries', {
        id: 0,
        entry: 'whats up everyone',
        date: '0',
        user_name: 'Lisa'
    })
    .then(function (response) {
        return response;
    })
    .catch(function (error) {
        return error;
    })
};

export default Post;