import { FacultiesInterface } from "./IFaculty";
import { OfficersInterface } from "./IOfficer";
import { TeachersInterface } from "./ITeacher";
import { TimesInterface } from "./ITime";

export interface SubjectsInterface {
    ID?: number;
    Code?: string;
    Name?: string;
    Credit?: number;
    Section?: number;
    Teacher?: TeachersInterface;
    TeacherID?: number;     // foreignkey.ID?
    Day?: string;
    Time?: TimesInterface; 
    TimeID?: number;     // foreignkey.ID?
    Take?: number;
    Faculty?: FacultiesInterface;
    FacultyID?: number; // foreignkey.ID?
    Officer?: OfficersInterface;
    OfficerID?: number;  // foreignkey.ID?
}