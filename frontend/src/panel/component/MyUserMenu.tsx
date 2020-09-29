import React from "react";

import { UserMenu } from "react-admin"
import MyAvatar from "./MyAvatar";


const MyUserMenu = props => (<UserMenu {...props} icon={<MyAvatar />} />);

export default MyUserMenu