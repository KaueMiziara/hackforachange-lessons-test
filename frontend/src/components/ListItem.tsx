import React from "react";
import "./ListItem.css";

const ListItem: React.FC<{name: string}> = ({name}) => {
    return (
        <div className="card">
            <p>{name}</p>
        </div>
    );
};

export default ListItem;
