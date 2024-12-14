import React, {useEffect, useState} from "react";
import {fetchPolls} from "../api";

const PollList = ({ onSelectPoll }) => {
    const [polls, setPolls] = useState([]);

    useEffect(() => {
        fetchPolls().then(setPolls);
    }, []);

    return (
        <div>
            <h2>Polls</h2>
            <ul>
             {polls.map((poll) => {
                <li key={poll.id} onClick={() => onSelectPoll(poll)}>
                    {poll.question}
                </li>
             })}
            </ul>
        </div>
    )
}