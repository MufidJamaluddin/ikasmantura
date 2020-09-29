import { Notification } from 'react-admin';
import React from "react";

const MyNotification = props => <Notification {...props} autoHideDuration={5000} />;

export default MyNotification;