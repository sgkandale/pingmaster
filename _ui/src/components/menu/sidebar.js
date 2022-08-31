import { Logout, Dashboard, InsertLink } from '@mui/icons-material';
import {
    Box, Drawer, List, ListItemText,
    Divider, ListItem, ListItemButton, ListItemIcon,
} from '@mui/material';
import { useNavigate } from 'react-router-dom';

const drawerWidth = 240;

export default function Sidebar(props) {
    const navigate = useNavigate()

    const menuItems = [
        {
            type: 'link',
            text: 'Overview',
            key: '/',
            icon: <Dashboard />,
            clickHandler: () => navigate('/')
        },
        {
            type: 'link',
            text: 'Targets',
            key: '/targets',
            icon: <InsertLink />,
            clickHandler: () => navigate('/targets')
        },
        {
            type: 'divider',
            text: 'divider-1'
        },
        {
            type: 'link',
            text: 'Logout',
            icon: <Logout />,
            clickHandler: props.callLogout
        }
    ]


    return <Drawer
        variant="permanent"
        sx={{
            width: drawerWidth,
            flexShrink: 0,
            [`& .MuiDrawer-paper`]: { width: drawerWidth, boxSizing: 'border-box' },
        }}
    >
        <Box sx={{ overflow: 'auto' }}>
            <List>
                {
                    menuItems.map((item) => {
                        if (item.type === 'divider') {
                            return <Divider key={item.text} />
                        }
                        return <ListItem
                            key={item.text}
                            disablePadding
                            onClick={item.clickHandler}
                            sx={{
                                bgcolor: window.location.pathname === item.key ?
                                    'primary.ultraLight' : 'white'
                            }}
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
}