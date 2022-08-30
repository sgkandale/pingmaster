import { Box, CssBaseline, Toolbar } from '@mui/material';
import { useNavigate } from 'react-router-dom';
import axios from 'axios'
import { ServerAddr } from '../server'
import { useDispatch, useSelector } from 'react-redux';
import { ACTION_LOGOUT } from '../state_actions'
import { useEffect, useState } from 'react';
import LogoutSnackbar from '../auth/logoutSnack';
import Topbar from './topbar';
import Sidebar from './sidebar';

const defaultSnackVal = {
    view: false,
    message: "",
    severity: "error"
}

export default function Menu(props) {
    const navigate = useNavigate();
    const user = useSelector(state => state.user)
    const loggedIn = useSelector(state => state.loggedIn)
    const dispatch = useDispatch()
    const [snack, setSnack] = useState(defaultSnackVal)

    const closeSnack = () => {
        setSnack(defaultSnackVal)
    }

    useEffect(() => {
        if (!loggedIn) {
            navigate('/login')
        }
    }, [loggedIn, navigate])


    const callLogout = () => {
        axios.post(ServerAddr + '/user/logout', {}, {
            headers: {
                'Content-Type': 'application/json',
                'Authorization': user.token
            }
        })
            .then(response => {
                dispatch({ type: ACTION_LOGOUT })
                navigate('/login')
            })
            .catch(error => {
                console.log(error)
                setSnack({
                    view: true,
                    message: error.response.data.message,
                    severity: "error"
                })
            });
    }

    // workaround to prevent rendering dashboard if not logged in
    if (!loggedIn) {
        return <></>
    }

    return <Box sx={{ display: 'flex' }}>
        <CssBaseline />
        <Topbar callLogout={callLogout} />
        <Sidebar callLogout={callLogout} />
        <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
            <Toolbar />
            {props.viewElement}
        </Box>
        <LogoutSnackbar open={snack.view} close={closeSnack} message={snack.message} />
    </Box>
}
