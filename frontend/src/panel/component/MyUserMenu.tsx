import React from "react";

import {UserMenu} from "react-admin"
import Avatar from '@material-ui/core/Avatar';


const MyUserMenu = props => (<UserMenu {...props} icon={<Avatar />} />);

export default MyUserMenu
