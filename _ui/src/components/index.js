import React from 'react';
import { Provider } from 'react-redux';
import { store } from './globalState';
import { CssBaseline } from '@mui/material'
import { ThemeProvider } from '@mui/material/styles'
import { theme1 } from './theme';
import Router from './router';

export default function Components() {
    
    return <Provider store={store}>
        <ThemeProvider theme={theme1}>
            <CssBaseline />
            <Router />
        </ThemeProvider>
    </Provider>
}