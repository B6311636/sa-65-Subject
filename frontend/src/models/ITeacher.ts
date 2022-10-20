import { FacultiesInterface } from "./IFaculty";

export interface TeachersInterface {
    ID: number,
    Level: string,
    Name: string,
    Email: string,
    // fk
    Faculty: FacultiesInterface,
    FacultyID: number
}