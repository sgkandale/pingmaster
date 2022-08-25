import {BrowserRouter,Routes,Route} from "react-router-dom";
import Auth from "./auth";

export default function Router() {
    return <BrowserRouter>
        <Routes>
            <Route path="/" element={<h1>/</h1>}>
            </Route>
            <Route 
                path="/login" 
                element={<Auth />}
            >
            </Route>
        </Routes>
    </BrowserRouter>
}