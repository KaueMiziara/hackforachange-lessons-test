import React from "react";

const ListItem: React.FC<{name: string}> = ({name}) => {
    return (
        <div>
            <p>{name}</p>
        </div>
    );
};

export default ListItem;
