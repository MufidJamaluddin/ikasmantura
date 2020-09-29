import { makeStyles } from '@material-ui/core/styles';
import Avatar from '@material-ui/core/Avatar';
import React from "react";

const useStyles = makeStyles({
    avatar: {
        height: 30,
        width: 30,
    },
});

const MyAvatar = () => {
    const classes = useStyles();
    return (
        <Avatar
            className={classes.avatar}
            src={process.env.PUBLIC_URL + '/static/img/ringga.jpg'}
        />
    )
};

export default MyAvatar