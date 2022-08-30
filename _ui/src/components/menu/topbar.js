import { Logout } from '@mui/icons-material';
import { Toolbar, Typography, IconButton, AppBar } from '@mui/material';

export default function Topbar(props) {
    return <AppBar
        position="fixed"
        sx={{ zIndex: (theme) => theme.zIndex.drawer + 1 }}
    >
        <Toolbar>
            <Typography
                variant="h5"
                sx={{
                    fontFamily: 'Silkscreen, Helvetica, sans-serif',
                    flexGrow: 1
                }}
            >
                pingmaster
            </Typography>
            <IconButton aria-label="logout" onClick={props.callLogout}>
                <Logout />
            </IconButton>
        </Toolbar>
    </AppBar>
}