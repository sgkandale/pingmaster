import { createTheme } from '@mui/material/styles'

export const theme1 = createTheme({
    palette: {
        primary: {
            ultraLight: '#CDEAFF',
            light: '#81CAFF',
            main: '#078DEE',
            dark: '#2381C5',
            ultraDark: '#1B659A',
        },
        secondary: {
            main: '#db4437',
        },
        background: {
            paper: '#FFFFFF',
            page: '#F5F9FD',
            dark: '#F5F5F5',
            darker: '#EBEBEB',
        },
        error: {
            ultraLight: '#FFC8C4',
            light: '#FF8A81',
            main: '#F44336',
        },
        success: {
            main: '#54B258',
        },
        text: {
            primary: '#202020',
            secondary: '#8B8B8B',
            disabled: '#CECECE',
        },
        info: {
            main: '#2196F3',
        },
    },
});