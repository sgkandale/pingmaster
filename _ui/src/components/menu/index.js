import {
    Box, Drawer, AppBar, CssBaseline, Toolbar, List, ListItemText,
    Typography, Divider, ListItem, ListItemButton, ListItemIcon, IconButton,
} from '@mui/material';
import { Dashboard, Logout } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';

const drawerWidth = 240;

export default function Menu(props) {
    const navigate = useNavigate();

    const menuItems = [
        {
            type: 'link',
            text: 'Overview',
            icon: <Dashboard />,
            clickHandler: () => navigate('/')
        },
        {
            type: 'divider'
        },
        {
            type: 'link',
            text: 'Logout',
            icon: <Logout />,
            clickHandler: () => { }
        }
    ]

    return (
        <Box sx={{ display: 'flex' }}>
            <CssBaseline />
            <AppBar
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
                    <IconButton aria-label="logout">
                        <Logout />
                    </IconButton>
                </Toolbar>
            </AppBar>
            <Drawer
                variant="permanent"
                sx={{
                    width: drawerWidth,
                    flexShrink: 0,
                    [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
                }}
            >
                <Toolbar />
                <Box sx={{ overflow: 'auto' }}>
                    <List>
                        {
                            menuItems.map((item) => {
                                if (item.type === 'divider') {
                                    return <Divider />
                                }
                                return <ListItem
                                    key={item.text}
                                    disablePadding
                                    onClick={item.clickHandler}
                                >
                                    <ListItemButton>
                                        <ListItemIcon>
                                            {item.icon}
                                        </ListItemIcon>
                                        <ListItemText primary={item.text} />
                                    </ListItemButton>
                                </ListItem>
                            })
                        }
                    </List>
                </Box>
            </Drawer>
            <Box component="main" sx={{ flexGrow: 1, p: 3 }}>
                <Toolbar />
                {props.viewElement}
            </Box>
        </Box>
    );
}
