import React, { useState, useEffect } from "react";
import { Link as RouterLink } from "react-router-dom";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Container from "@mui/material/Container";
import Box from "@mui/material/Box";
import { OfficersInterface } from "../models/IOfficer";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
function Officers() {
    const [officers, setOfficers] = useState<OfficersInterface[]>([]);

    const apiUrl = "http://localhost:8080/officers";

    useEffect(() => {
        getOfficers();
    }, []);

    const getOfficers = async () => {
        const requestOptions = {
            method: "GET",
            headers: {
                Authorization: `Bearer ${localStorage.getItem("token")}`,
                "Content-Type": "application/json",
            },
        };
    
        fetch(apiUrl, requestOptions)
            .then((response) => response.json())
            .then((res) => {
                if (res.data) {
                    return setOfficers(res.data);
                } else {
                    return false;
                }
            });
    };

    const columns: GridColDef[] = [
        { field: "ID", headerName: "ลำดับ", width: 100 },
        { field: "Name", headerName: "ชื่อ - สกุล", width: 300 },
        { field: "Email", headerName: "อีเมล", width: 400 },
    ];

    return (
        <div>
            <Container maxWidth="md">
                <Box display="flex" sx={{marginTop: 2,}}>
                    <Box flexGrow={1}>
                        <Typography component="h2" variant="h6" color="primary" gutterBottom>
                            ข้อมูลสมาชิกผู้ดูแลระบบ
                        </Typography>
                    </Box>
                    <Box>
                        <Button component={RouterLink} to="/officer/create" variant="contained" color="primary">
                            สร้างข้อมูล
                        </Button>
                    </Box>
                </Box>
                <div style={{ height: 400, width: "100%", marginTop: "20px" }}>
                    <DataGrid
                        rows={officers}
                        getRowId={(row) => row.ID}
                        columns={columns}
                        pageSize={5}
                        rowsPerPageOptions={[5]}
                    />
                </div>
            </Container>
        </div>
    );
}

export default Officers;