package hackerrank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type ATSCodePairCreate struct {
	Candidate         CandidateInformation `json:"candidate"`
	CandidateId       string               `json:"candidate_id"`
	InterviewMetadata InterviewMetadata    `json:"interview_metadata"`
	RequisitionId     string               `json:"requisition_id"`
	SendEmail         bool                 `json:"send_email"`
	Title             string               `json:"title"`
}
type ATSCodePairIndex struct {
	CandidateId   string `json:"candidate_id"`
	RequisitionId string `json:"requisition_id"`
	Title         string `json:"title"`
}
type ATSCodePairRef struct {
	Candidate         map[string]interface{} `json:"candidate"`
	CandidateId       string                 `json:"candidate_id"`
	InterviewMetadata map[string]interface{} `json:"interview_metadata"`
	RequisitionId     string                 `json:"requisition_id"`
	SendEmail         bool                   `json:"send_email"`
	Title             string                 `json:"title"`
}
type ATSCodePairShow struct {
	CandidateId   string `json:"candidate_id"`
	RequisitionId string `json:"requisition_id"`
	Title         string `json:"title"`
}
type ATSCodePairUpdate struct {
	CandidateId   string `json:"candidate_id"`
	RequisitionId string `json:"requisition_id"`
	Title         string `json:"title"`
}
type ATSCodeScreenCreate struct {
	AcceptResultUpdates bool                    `json:"accept_result_updates"`
	Accommodations      CandidateAccommodations `json:"accommodations"`
	CandidateId         string                  `json:"candidate_id"`
	Email               string                  `json:"email"`
	Force               bool                    `json:"force"`
	ForceReattemptAfter float32                 `json:"force_reattempt_after"`
	RequisitionId       string                  `json:"requisition_id"`
	SendEmail           bool                    `json:"send_email"`
	TestId              string                  `json:"test_id"`
	TestResultUrl       string                  `json:"test_result_url"`
}
type ATSCodeScreenIndex struct {
	AcceptResultUpdates bool   `json:"accept_result_updates"`
	CandidateId         string `json:"candidate_id"`
	Email               string `json:"email"`
	RequisitionId       string `json:"requisition_id"`
	TestId              string `json:"test_id"`
	TestResultUrl       string `json:"test_result_url"`
}
type ATSCodeScreenRef struct {
	AcceptResultUpdates bool                   `json:"accept_result_updates"`
	Accommodations      map[string]interface{} `json:"accommodations"`
	CandidateId         string                 `json:"candidate_id"`
	Email               string                 `json:"email"`
	Force               bool                   `json:"force"`
	ForceReattemptAfter float32                `json:"force_reattempt_after"`
	RequisitionId       string                 `json:"requisition_id"`
	SendEmail           bool                   `json:"send_email"`
	TestId              string                 `json:"test_id"`
	TestResultUrl       string                 `json:"test_result_url"`
}
type ATSCodeScreenShow struct {
	AcceptResultUpdates bool   `json:"accept_result_updates"`
	CandidateId         string `json:"candidate_id"`
	Email               string `json:"email"`
	RequisitionId       string `json:"requisition_id"`
	TestId              string `json:"test_id"`
	TestResultUrl       string `json:"test_result_url"`
}
type ATSCodeScreenUpdate struct {
	AcceptResultUpdates bool   `json:"accept_result_updates"`
	CandidateId         string `json:"candidate_id"`
	RequisitionId       string `json:"requisition_id"`
	TestId              string `json:"test_id"`
	TestResultUrl       string `json:"test_result_url"`
}
type AuditLogCreate struct {
}
type AuditLogIndex struct {
	Action         string                 `json:"action"`
	CreatedAt      string                 `json:"created_at"`
	IpAddress      string                 `json:"ip_address"`
	ModifiedFields []string               `json:"modified_fields"`
	ModifiedValues map[string]interface{} `json:"modified_values"`
	SourceId       float32                `json:"source_id"`
	SourceType     string                 `json:"source_type"`
	User           string                 `json:"user"`
}
type AuditLogRef struct {
	Action         string                 `json:"action"`
	CreatedAt      string                 `json:"created_at"`
	IpAddress      string                 `json:"ip_address"`
	ModifiedFields []string               `json:"modified_fields"`
	ModifiedValues map[string]interface{} `json:"modified_values"`
	SourceId       float32                `json:"source_id"`
	SourceType     string                 `json:"source_type"`
}
type AuditLogShow struct {
	Action         string                 `json:"action"`
	CreatedAt      string                 `json:"created_at"`
	IpAddress      string                 `json:"ip_address"`
	ModifiedFields []string               `json:"modified_fields"`
	ModifiedValues map[string]interface{} `json:"modified_values"`
	SourceId       float32                `json:"source_id"`
	SourceType     string                 `json:"source_type"`
	User           string                 `json:"user"`
}
type AuditLogUpdate struct {
}
type CandidateAccommodations struct {
	AdditionalTimePercent float32 `json:"additional_time_percent"`
}
type CandidateDetail struct {
	FieldName string `json:"field_name"`
	Title     string `json:"title"`
	Value     string `json:"value"`
}
type CandidateInformation struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type InterviewCreate struct {
	Candidate           CandidateInformation   `json:"candidate"`
	From                string                 `json:"from"`
	InterviewTemplateId float32                `json:"interview_template_id"`
	Interviewers        []string               `json:"interviewers"`
	Metadata            map[string]interface{} `json:"metadata"`
	Notes               string                 `json:"notes"`
	ResultUrl           string                 `json:"result_url"`
	ResumeUrl           string                 `json:"resume_url"`
	SendEmail           bool                   `json:"send_email"`
	Title               string                 `json:"title"`
	To                  string                 `json:"to"`
}
type InterviewIndex struct {
	Candidate           map[string]interface{} `json:"candidate"`
	CreatedAt           string                 `json:"created_at"`
	EndedAt             string                 `json:"ended_at"`
	Feedback            string                 `json:"feedback"`
	From                string                 `json:"from"`
	Id                  string                 `json:"id"`
	InterviewTemplateId float32                `json:"interview_template_id"`
	Interviewers        []string               `json:"interviewers"`
	Metadata            map[string]interface{} `json:"metadata"`
	Notes               string                 `json:"notes"`
	ReportUrl           string                 `json:"report_url"`
	ResultUrl           string                 `json:"result_url"`
	ResumeUrl           string                 `json:"resume_url"`
	Status              string                 `json:"status"`
	ThumbsUp            float32                `json:"thumbs_up"`
	Title               string                 `json:"title"`
	To                  string                 `json:"to"`
	UpdatedAt           string                 `json:"updated_at"`
	Url                 string                 `json:"url"`
	User                float32                `json:"user"`
}
type InterviewMetadata struct {
	AwsRequestId string `json:"aws_request_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
}
type InterviewRef struct {
	Candidate           map[string]interface{} `json:"candidate"`
	CreatedAt           string                 `json:"created_at"`
	EndedAt             string                 `json:"ended_at"`
	Feedback            string                 `json:"feedback"`
	From                string                 `json:"from"`
	Id                  string                 `json:"id"`
	InterviewTemplateId float32                `json:"interview_template_id"`
	Interviewers        []string               `json:"interviewers"`
	Metadata            map[string]interface{} `json:"metadata"`
	Notes               string                 `json:"notes"`
	ReportUrl           string                 `json:"report_url"`
	ResultUrl           string                 `json:"result_url"`
	ResumeUrl           string                 `json:"resume_url"`
	SendEmail           bool                   `json:"send_email"`
	Status              string                 `json:"status"`
	ThumbsUp            float32                `json:"thumbs_up"`
	Title               string                 `json:"title"`
	To                  string                 `json:"to"`
	UpdatedAt           string                 `json:"updated_at"`
	Url                 string                 `json:"url"`
}
type InterviewShow struct {
	Candidate           map[string]interface{} `json:"candidate"`
	CreatedAt           string                 `json:"created_at"`
	EndedAt             string                 `json:"ended_at"`
	Feedback            string                 `json:"feedback"`
	From                string                 `json:"from"`
	Id                  string                 `json:"id"`
	InterviewTemplateId float32                `json:"interview_template_id"`
	Interviewers        []string               `json:"interviewers"`
	Metadata            map[string]interface{} `json:"metadata"`
	Notes               string                 `json:"notes"`
	ReportUrl           string                 `json:"report_url"`
	ResultUrl           string                 `json:"result_url"`
	ResumeUrl           string                 `json:"resume_url"`
	Status              string                 `json:"status"`
	ThumbsUp            float32                `json:"thumbs_up"`
	Title               string                 `json:"title"`
	To                  string                 `json:"to"`
	UpdatedAt           string                 `json:"updated_at"`
	Url                 string                 `json:"url"`
	User                float32                `json:"user"`
}
type InterviewUpdate struct {
	Candidate           map[string]interface{} `json:"candidate"`
	From                string                 `json:"from"`
	InterviewTemplateId float32                `json:"interview_template_id"`
	Metadata            map[string]interface{} `json:"metadata"`
	Notes               string                 `json:"notes"`
	ResultUrl           string                 `json:"result_url"`
	ResumeUrl           string                 `json:"resume_url"`
	SendEmail           bool                   `json:"send_email"`
	Title               string                 `json:"title"`
	To                  string                 `json:"to"`
}
type InvitersCreate struct {
	CandidatesInvited    float32          `json:"candidates_invited"`
	CandidatesPermission float32          `json:"candidates_permission"`
	Email                string           `json:"email"`
	Firstname            string           `json:"firstname"`
	InterviewsPermission float32          `json:"interviews_permission"`
	Lastname             string           `json:"lastname"`
	Phone                string           `json:"phone"`
	QuestionsPermission  float32          `json:"questions_permission"`
	SendEmail            bool             `json:"send_email"`
	Teams                []UserTeamCreate `json:"teams"`
	TestsPermission      float32          `json:"tests_permission"`
}
type InvitersIndex struct {
	Activated            bool     `json:"activated"`
	CandidatesInvited    float32  `json:"candidates_invited"`
	CandidatesPermission float32  `json:"candidates_permission"`
	Email                string   `json:"email"`
	Firstname            string   `json:"firstname"`
	Id                   string   `json:"id"`
	InterviewsPermission float32  `json:"interviews_permission"`
	Lastname             string   `json:"lastname"`
	Phone                string   `json:"phone"`
	QuestionsPermission  float32  `json:"questions_permission"`
	Role                 string   `json:"role"`
	Status               string   `json:"status"`
	Teams                []string `json:"teams"`
	TestsPermission      float32  `json:"tests_permission"`
	Timezone             string   `json:"timezone"`
}
type InvitersRef struct {
	Activated            bool    `json:"activated"`
	CandidatesInvited    float32 `json:"candidates_invited"`
	CandidatesPermission float32 `json:"candidates_permission"`
	Email                string  `json:"email"`
	Firstname            string  `json:"firstname"`
	Id                   string  `json:"id"`
	InterviewsPermission float32 `json:"interviews_permission"`
	Lastname             string  `json:"lastname"`
	Phone                string  `json:"phone"`
	QuestionsPermission  float32 `json:"questions_permission"`
	Role                 string  `json:"role"`
	SendEmail            bool    `json:"send_email"`
	Status               string  `json:"status"`
	TestsPermission      float32 `json:"tests_permission"`
	Timezone             string  `json:"timezone"`
}
type InvitersShow struct {
	Activated            bool     `json:"activated"`
	CandidatesInvited    float32  `json:"candidates_invited"`
	CandidatesPermission float32  `json:"candidates_permission"`
	Email                string   `json:"email"`
	Firstname            string   `json:"firstname"`
	Id                   string   `json:"id"`
	InterviewsPermission float32  `json:"interviews_permission"`
	Lastname             string   `json:"lastname"`
	Phone                string   `json:"phone"`
	QuestionsPermission  float32  `json:"questions_permission"`
	Role                 string   `json:"role"`
	Status               string   `json:"status"`
	Teams                []string `json:"teams"`
	TestsPermission      float32  `json:"tests_permission"`
	Timezone             string   `json:"timezone"`
}
type InvitersUpdate struct {
	CandidatesInvited    float32 `json:"candidates_invited"`
	CandidatesPermission float32 `json:"candidates_permission"`
	Firstname            string  `json:"firstname"`
	InterviewsPermission float32 `json:"interviews_permission"`
	Lastname             string  `json:"lastname"`
	Phone                string  `json:"phone"`
	QuestionsPermission  float32 `json:"questions_permission"`
	TestsPermission      float32 `json:"tests_permission"`
}
type PatchOperationObject struct {
	Op    string                 `json:"op"`
	Value map[string]interface{} `json:"value"`
}
type QuestionCreate struct {
	InternalNotes       string   `json:"internal_notes"`
	Languages           []string `json:"languages"`
	Name                string   `json:"name"`
	ProblemStatement    string   `json:"problem_statement"`
	RecommendedDuration float32  `json:"recommended_duration"`
	Tags                []string `json:"tags"`
	Type                string   `json:"type"`
}
type QuestionIndex struct {
	CreatedAt           string   `json:"created_at"`
	Id                  string   `json:"id"`
	InternalNotes       string   `json:"internal_notes"`
	Languages           []string `json:"languages"`
	MaxScore            float32  `json:"max_score"`
	Name                string   `json:"name"`
	Owner               string   `json:"owner"`
	ProblemStatement    string   `json:"problem_statement"`
	RecommendedDuration float32  `json:"recommended_duration"`
	Status              string   `json:"status"`
	Tags                []string `json:"tags"`
	Type                string   `json:"type"`
	UniqueId            string   `json:"unique_id"`
}
type QuestionRef struct {
	CreatedAt           string   `json:"created_at"`
	Id                  string   `json:"id"`
	InternalNotes       string   `json:"internal_notes"`
	Languages           []string `json:"languages"`
	MaxScore            float32  `json:"max_score"`
	Name                string   `json:"name"`
	Owner               string   `json:"owner"`
	ProblemStatement    string   `json:"problem_statement"`
	RecommendedDuration float32  `json:"recommended_duration"`
	Status              string   `json:"status"`
	Tags                []string `json:"tags"`
	Type                string   `json:"type"`
	UniqueId            string   `json:"unique_id"`
}
type QuestionScore struct {
	Answer   string  `json:"answer"`
	Answered bool    `json:"answered"`
	Name     string  `json:"name"`
	Preview  string  `json:"preview"`
	Score    float32 `json:"score"`
}
type QuestionShow struct {
	CreatedAt           string   `json:"created_at"`
	Id                  string   `json:"id"`
	InternalNotes       string   `json:"internal_notes"`
	Languages           []string `json:"languages"`
	MaxScore            float32  `json:"max_score"`
	Name                string   `json:"name"`
	Owner               string   `json:"owner"`
	ProblemStatement    string   `json:"problem_statement"`
	RecommendedDuration float32  `json:"recommended_duration"`
	Status              string   `json:"status"`
	Tags                []string `json:"tags"`
	Type                string   `json:"type"`
	UniqueId            string   `json:"unique_id"`
}
type QuestionUpdate struct {
	InternalNotes       string   `json:"internal_notes"`
	Languages           []string `json:"languages"`
	Name                string   `json:"name"`
	ProblemStatement    string   `json:"problem_statement"`
	RecommendedDuration float32  `json:"recommended_duration"`
	Tags                []string `json:"tags"`
	Type                string   `json:"type"`
}
type SCIMPatchCreate struct {
	Operations []PatchOperationObject `json:"operations"`
}
type SCIMPatchIndex struct {
	Operations []string `json:"operations"`
}
type SCIMPatchRef struct {
	Operations []string `json:"operations"`
}
type SCIMPatchShow struct {
	Operations []string `json:"operations"`
}
type SCIMPatchUpdate struct {
	Operations []string `json:"operations"`
}
type SCIMTeamCreate struct {
	DiplayName string `json:"diplayName"`
}
type SCIMTeamEntryCreate struct {
	DiplayName string `json:"diplayName"`
}
type SCIMTeamEntryIndex struct {
	DiplayName string `json:"diplayName"`
	Id         string `json:"id"`
}
type SCIMTeamEntryRef struct {
	DiplayName string `json:"diplayName"`
	Id         string `json:"id"`
}
type SCIMTeamEntryShow struct {
	DiplayName string `json:"diplayName"`
	Id         string `json:"id"`
}
type SCIMTeamEntryUpdate struct {
	DiplayName string `json:"diplayName"`
}
type SCIMTeamIndex struct {
	DiplayName string   `json:"diplayName"`
	Id         string   `json:"id"`
	Schemas    []string `json:"schemas"`
}
type SCIMTeamRef struct {
	DiplayName string   `json:"diplayName"`
	Id         string   `json:"id"`
	Schemas    []string `json:"schemas"`
}
type SCIMTeamShow struct {
	DiplayName string   `json:"diplayName"`
	Id         string   `json:"id"`
	Schemas    []string `json:"schemas"`
}
type SCIMTeamUpdate struct {
	DiplayName string `json:"diplayName"`
}
type SCIMUserCreate struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserEntryCreate struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserEntryIndex struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Id           string                 `json:"id"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserEntryRef struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Id           string                 `json:"id"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserEntryShow struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Id           string                 `json:"id"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserEntryUpdate struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
}
type SCIMUserIndex struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Id           string                 `json:"id"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	Schemas      []string               `json:"schemas"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserRef struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Id           string                 `json:"id"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	Schemas      []string               `json:"schemas"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserShow struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Id           string                 `json:"id"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	Schemas      []string               `json:"schemas"`
	TeamAdmin    bool                   `json:"team_admin"`
	UserName     string                 `json:"userName"`
}
type SCIMUserUpdate struct {
	Active       bool                   `json:"active"`
	CompanyAdmin bool                   `json:"company_admin"`
	Emails       []string               `json:"emails"`
	Name         map[string]interface{} `json:"name"`
	Role         string                 `json:"role"`
	TeamAdmin    bool                   `json:"team_admin"`
}
type TeamCreate struct {
	Departments  []string `json:"departments"`
	DeveloperCap float32  `json:"developer_cap"`
	InviteAs     string   `json:"invite_as"`
	Locations    []string `json:"locations"`
	Name         string   `json:"name"`
	RecruiterCap float32  `json:"recruiter_cap"`
}
type TeamIndex struct {
	CreatedAt      string   `json:"created_at"`
	Departments    []string `json:"departments"`
	DeveloperCap   float32  `json:"developer_cap"`
	DeveloperCount float32  `json:"developer_count"`
	Id             string   `json:"id"`
	InviteAs       string   `json:"invite_as"`
	Locations      []string `json:"locations"`
	Name           string   `json:"name"`
	Owner          string   `json:"owner"`
	RecruiterCap   float32  `json:"recruiter_cap"`
	RecruiterCount float32  `json:"recruiter_count"`
}
type TeamRef struct {
	CreatedAt      string   `json:"created_at"`
	Departments    []string `json:"departments"`
	DeveloperCap   float32  `json:"developer_cap"`
	DeveloperCount float32  `json:"developer_count"`
	Id             string   `json:"id"`
	InviteAs       string   `json:"invite_as"`
	Locations      []string `json:"locations"`
	Name           string   `json:"name"`
	RecruiterCap   float32  `json:"recruiter_cap"`
	RecruiterCount float32  `json:"recruiter_count"`
}
type TeamShow struct {
	CreatedAt      string   `json:"created_at"`
	Departments    []string `json:"departments"`
	DeveloperCap   float32  `json:"developer_cap"`
	DeveloperCount float32  `json:"developer_count"`
	Id             string   `json:"id"`
	InviteAs       string   `json:"invite_as"`
	Locations      []string `json:"locations"`
	Name           string   `json:"name"`
	Owner          string   `json:"owner"`
	RecruiterCap   float32  `json:"recruiter_cap"`
	RecruiterCount float32  `json:"recruiter_count"`
}
type TeamUpdate struct {
	Departments  []string `json:"departments"`
	DeveloperCap float32  `json:"developer_cap"`
	InviteAs     string   `json:"invite_as"`
	Locations    []string `json:"locations"`
	Name         string   `json:"name"`
	RecruiterCap float32  `json:"recruiter_cap"`
}
type TemplateCreate struct {
	Content string `json:"content"`
	Default bool   `json:"default"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}
type TemplateIndex struct {
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Default   bool   `json:"default"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Subject   string `json:"subject"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
}
type TemplateRef struct {
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Default   bool   `json:"default"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Subject   string `json:"subject"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
}
type TemplateShow struct {
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	Default   bool   `json:"default"`
	Id        string `json:"id"`
	Name      string `json:"name"`
	Subject   string `json:"subject"`
	UpdatedAt string `json:"updated_at"`
	User      string `json:"user"`
}
type TemplateUpdate struct {
	Content string `json:"content"`
	Default bool   `json:"default"`
	Name    string `json:"name"`
	Subject string `json:"subject"`
}
type TestCandidateCreate struct {
	AcceptResultUpdates bool                    `json:"accept_result_updates"`
	Accommodations      CandidateAccommodations `json:"accommodations"`
	AtsState            float32                 `json:"ats_state"`
	Email               string                  `json:"email"`
	EvaluatorEmail      string                  `json:"evaluator_email"`
	Force               bool                    `json:"force"`
	ForceReattempt      bool                    `json:"force_reattempt"`
	FullName            string                  `json:"full_name"`
	InviteMetadata      map[string]interface{}  `json:"invite_metadata"`
	InviteValidFrom     string                  `json:"invite_valid_from"`
	InviteValidTo       string                  `json:"invite_valid_to"`
	Message             string                  `json:"message"`
	SendEmail           bool                    `json:"send_email"`
	Subject             string                  `json:"subject"`
	Tags                []string                `json:"tags"`
	Template            string                  `json:"template"`
	TestFinishUrl       string                  `json:"test_finish_url"`
	TestResultUrl       string                  `json:"test_result_url"`
}
type TestCandidateIndex struct {
	AcceptResultUpdates    bool                   `json:"accept_result_updates"`
	AddedTime              float32                `json:"added_time"`
	AtsState               float32                `json:"ats_state"`
	AttemptEndtime         string                 `json:"attempt_endtime"`
	AttemptStarttime       string                 `json:"attempt_starttime"`
	AuthenticatedReportUrl string                 `json:"authenticated_report_url"`
	CandidateDetails       []CandidateDetail      `json:"candidate_details"`
	Comments               map[string]interface{} `json:"comments"`
	EditorPasteCount       float32                `json:"editor_paste_count"`
	Email                  string                 `json:"email"`
	EvaluatorEmail         string                 `json:"evaluator_email"`
	Feedback               string                 `json:"feedback"`
	FullName               string                 `json:"full_name"`
	Id                     string                 `json:"id"`
	InviteEmailDone        bool                   `json:"invite_email_done"`
	InviteLink             string                 `json:"invite_link"`
	InviteMetadata         map[string]interface{} `json:"invite_metadata"`
	InviteValid            bool                   `json:"invite_valid"`
	InviteValidFrom        string                 `json:"invite_valid_from"`
	InviteValidTo          string                 `json:"invite_valid_to"`
	InvitedOn              string                 `json:"invited_on"`
	MaxCodeSimilarity      map[string]interface{} `json:"max_code_similarity"`
	OutOfWindowDuration    float32                `json:"out_of_window_duration"`
	OutOfWindowEvents      float32                `json:"out_of_window_events"`
	PdfUrl                 string                 `json:"pdf_url"`
	PercentageScore        float32                `json:"percentage_score"`
	Plagiarism             map[string]interface{} `json:"plagiarism"`
	PlagiarismStatus       bool                   `json:"plagiarism_status"`
	ProctorImages          []string               `json:"proctor_images"`
	Questions              QuestionScore          `json:"questions"`
	ReportUrl              string                 `json:"report_url"`
	Score                  float32                `json:"score"`
	ScoresTagsSplit        map[string]interface{} `json:"scores_tags_split"`
	Status                 float32                `json:"status"`
	Tags                   []string               `json:"tags"`
	Test                   string                 `json:"test"`
	TestFinishUrl          string                 `json:"test_finish_url"`
	TestResultUrl          string                 `json:"test_result_url"`
	UnclaimedAddedTime     float32                `json:"unclaimed_added_time"`
	User                   string                 `json:"user"`
}
type TestCandidateRef struct {
	AcceptResultUpdates    bool                   `json:"accept_result_updates"`
	Accommodations         map[string]interface{} `json:"accommodations"`
	AddedTime              float32                `json:"added_time"`
	AtsState               float32                `json:"ats_state"`
	AttemptEndtime         string                 `json:"attempt_endtime"`
	AttemptStarttime       string                 `json:"attempt_starttime"`
	AuthenticatedReportUrl string                 `json:"authenticated_report_url"`
	CandidateDetails       []CandidateDetail      `json:"candidate_details"`
	Comments               map[string]interface{} `json:"comments"`
	EditorPasteCount       float32                `json:"editor_paste_count"`
	Email                  string                 `json:"email"`
	EvaluatorEmail         string                 `json:"evaluator_email"`
	Feedback               string                 `json:"feedback"`
	Force                  bool                   `json:"force"`
	ForceReattempt         bool                   `json:"force_reattempt"`
	FullName               string                 `json:"full_name"`
	Id                     string                 `json:"id"`
	InviteEmailDone        bool                   `json:"invite_email_done"`
	InviteLink             string                 `json:"invite_link"`
	InviteMetadata         map[string]interface{} `json:"invite_metadata"`
	InviteValid            bool                   `json:"invite_valid"`
	InviteValidFrom        string                 `json:"invite_valid_from"`
	InviteValidTo          string                 `json:"invite_valid_to"`
	InvitedOn              string                 `json:"invited_on"`
	MaxCodeSimilarity      map[string]interface{} `json:"max_code_similarity"`
	Message                string                 `json:"message"`
	OutOfWindowDuration    float32                `json:"out_of_window_duration"`
	OutOfWindowEvents      float32                `json:"out_of_window_events"`
	PdfUrl                 string                 `json:"pdf_url"`
	PercentageScore        float32                `json:"percentage_score"`
	Plagiarism             map[string]interface{} `json:"plagiarism"`
	PlagiarismStatus       bool                   `json:"plagiarism_status"`
	ProctorImages          []string               `json:"proctor_images"`
	ReportUrl              string                 `json:"report_url"`
	Score                  float32                `json:"score"`
	ScoresTagsSplit        map[string]interface{} `json:"scores_tags_split"`
	SendEmail              bool                   `json:"send_email"`
	Status                 float32                `json:"status"`
	Subject                string                 `json:"subject"`
	Tags                   []string               `json:"tags"`
	Template               string                 `json:"template"`
	TestFinishUrl          string                 `json:"test_finish_url"`
	TestResultUrl          string                 `json:"test_result_url"`
	UnclaimedAddedTime     float32                `json:"unclaimed_added_time"`
}
type TestCandidateShow struct {
	AcceptResultUpdates    bool                   `json:"accept_result_updates"`
	AddedTime              float32                `json:"added_time"`
	AtsState               float32                `json:"ats_state"`
	AttemptEndtime         string                 `json:"attempt_endtime"`
	AttemptStarttime       string                 `json:"attempt_starttime"`
	AuthenticatedReportUrl string                 `json:"authenticated_report_url"`
	CandidateDetails       []CandidateDetail      `json:"candidate_details"`
	Comments               map[string]interface{} `json:"comments"`
	EditorPasteCount       float32                `json:"editor_paste_count"`
	Email                  string                 `json:"email"`
	EvaluatorEmail         string                 `json:"evaluator_email"`
	Feedback               string                 `json:"feedback"`
	FullName               string                 `json:"full_name"`
	Id                     string                 `json:"id"`
	InviteEmailDone        bool                   `json:"invite_email_done"`
	InviteLink             string                 `json:"invite_link"`
	InviteMetadata         map[string]interface{} `json:"invite_metadata"`
	InviteValid            bool                   `json:"invite_valid"`
	InviteValidFrom        string                 `json:"invite_valid_from"`
	InviteValidTo          string                 `json:"invite_valid_to"`
	InvitedOn              string                 `json:"invited_on"`
	MaxCodeSimilarity      map[string]interface{} `json:"max_code_similarity"`
	OutOfWindowDuration    float32                `json:"out_of_window_duration"`
	OutOfWindowEvents      float32                `json:"out_of_window_events"`
	PdfUrl                 string                 `json:"pdf_url"`
	PercentageScore        float32                `json:"percentage_score"`
	Plagiarism             map[string]interface{} `json:"plagiarism"`
	PlagiarismStatus       bool                   `json:"plagiarism_status"`
	ProctorImages          []string               `json:"proctor_images"`
	Questions              QuestionScore          `json:"questions"`
	ReportUrl              string                 `json:"report_url"`
	Score                  float32                `json:"score"`
	ScoresTagsSplit        map[string]interface{} `json:"scores_tags_split"`
	Status                 float32                `json:"status"`
	Tags                   []string               `json:"tags"`
	Test                   string                 `json:"test"`
	TestFinishUrl          string                 `json:"test_finish_url"`
	TestResultUrl          string                 `json:"test_result_url"`
	UnclaimedAddedTime     float32                `json:"unclaimed_added_time"`
	User                   string                 `json:"user"`
}
type TestCandidateUpdate struct {
	AcceptResultUpdates bool                   `json:"accept_result_updates"`
	Accommodations      map[string]interface{} `json:"accommodations"`
	AtsState            float32                `json:"ats_state"`
	EvaluatorEmail      string                 `json:"evaluator_email"`
	FullName            string                 `json:"full_name"`
	InviteMetadata      map[string]interface{} `json:"invite_metadata"`
	InviteValidFrom     string                 `json:"invite_valid_from"`
	InviteValidTo       string                 `json:"invite_valid_to"`
	Tags                []string               `json:"tags"`
	TestFinishUrl       string                 `json:"test_finish_url"`
	TestResultUrl       string                 `json:"test_result_url"`
}
type TestsCreate struct {
	CandidateDetails      []string `json:"candidate_details"`
	CandidateTabSwitch    bool     `json:"candidate_tab_switch"`
	CustomAcknowledgeText string   `json:"custom_acknowledge_text"`
	CutoffScore           float32  `json:"cutoff_score"`
	Draft                 bool     `json:"draft"`
	Duration              float32  `json:"duration"`
	EnableAcknowledgement bool     `json:"enable_acknowledgement"`
	EnableProctoring      bool     `json:"enable_proctoring"`
	Endtime               string   `json:"endtime"`
	Experience            []string `json:"experience"`
	HideCompileTest       bool     `json:"hide_compile_test"`
	HideTemplate          bool     `json:"hide_template"`
	IdeConfig             string   `json:"ide_config"`
	Instructions          string   `json:"instructions"`
	Languages             []string `json:"languages"`
	Locked                bool     `json:"locked"`
	MasterPassword        string   `json:"master_password"`
	McqCorrectScore       float32  `json:"mcq_correct_score"`
	McqIncorrectScore     float32  `json:"mcq_incorrect_score"`
	Name                  string   `json:"name"`
	Questions             []string `json:"questions"`
	RoleIds               []string `json:"role_ids"`
	Secure                bool     `json:"secure"`
	ShowCopyPastePrompt   bool     `json:"show_copy_paste_prompt"`
	ShuffleQuestions      bool     `json:"shuffle_questions"`
	Starttime             string   `json:"starttime"`
	Tags                  []string `json:"tags"`
	TestAdmins            []string `json:"test_admins"`
	TrackEditorPaste      bool     `json:"track_editor_paste"`
}
type TestsIndex struct {
	CandidateDetails      []string               `json:"candidate_details"`
	CandidateTabSwitch    bool                   `json:"candidate_tab_switch"`
	CreatedAt             string                 `json:"created_at"`
	CustomAcknowledgeText string                 `json:"custom_acknowledge_text"`
	CutoffScore           float32                `json:"cutoff_score"`
	Draft                 bool                   `json:"draft"`
	Duration              float32                `json:"duration"`
	EnableAcknowledgement bool                   `json:"enable_acknowledgement"`
	EnableProctoring      bool                   `json:"enable_proctoring"`
	Endtime               string                 `json:"endtime"`
	Experience            []string               `json:"experience"`
	HideCompileTest       bool                   `json:"hide_compile_test"`
	HideTemplate          bool                   `json:"hide_template"`
	Id                    string                 `json:"id"`
	IdeConfig             string                 `json:"ide_config"`
	Instructions          string                 `json:"instructions"`
	Languages             []string               `json:"languages"`
	Locked                bool                   `json:"locked"`
	LockedBy              string                 `json:"locked_by"`
	MasterPassword        string                 `json:"master_password"`
	McqCorrectScore       float32                `json:"mcq_correct_score"`
	McqIncorrectScore     float32                `json:"mcq_incorrect_score"`
	Name                  string                 `json:"name"`
	Owner                 string                 `json:"owner"`
	PublicLoginUrl        string                 `json:"public_login_url"`
	Questions             []string               `json:"questions"`
	RoleIds               []string               `json:"role_ids"`
	Sections              map[string]interface{} `json:"sections"`
	Secure                bool                   `json:"secure"`
	ShortLoginUrl         string                 `json:"short_login_url"`
	ShowCopyPastePrompt   bool                   `json:"show_copy_paste_prompt"`
	ShuffleQuestions      bool                   `json:"shuffle_questions"`
	Starred               bool                   `json:"starred"`
	Starttime             string                 `json:"starttime"`
	State                 string                 `json:"state"`
	Tags                  []string               `json:"tags"`
	TestAdmins            []string               `json:"test_admins"`
	TrackEditorPaste      bool                   `json:"track_editor_paste"`
	UniqueId              string                 `json:"unique_id"`
}
type TestsRef struct {
	CandidateDetails      []string               `json:"candidate_details"`
	CandidateTabSwitch    bool                   `json:"candidate_tab_switch"`
	CreatedAt             string                 `json:"created_at"`
	CustomAcknowledgeText string                 `json:"custom_acknowledge_text"`
	CutoffScore           float32                `json:"cutoff_score"`
	Draft                 bool                   `json:"draft"`
	Duration              float32                `json:"duration"`
	EnableAcknowledgement bool                   `json:"enable_acknowledgement"`
	EnableProctoring      bool                   `json:"enable_proctoring"`
	Endtime               string                 `json:"endtime"`
	Experience            []string               `json:"experience"`
	HideCompileTest       bool                   `json:"hide_compile_test"`
	HideTemplate          bool                   `json:"hide_template"`
	Id                    string                 `json:"id"`
	IdeConfig             string                 `json:"ide_config"`
	Instructions          string                 `json:"instructions"`
	Languages             []string               `json:"languages"`
	Locked                bool                   `json:"locked"`
	MasterPassword        string                 `json:"master_password"`
	McqCorrectScore       float32                `json:"mcq_correct_score"`
	McqIncorrectScore     float32                `json:"mcq_incorrect_score"`
	Name                  string                 `json:"name"`
	PublicLoginUrl        string                 `json:"public_login_url"`
	Questions             []string               `json:"questions"`
	RoleIds               []string               `json:"role_ids"`
	Sections              map[string]interface{} `json:"sections"`
	Secure                bool                   `json:"secure"`
	ShortLoginUrl         string                 `json:"short_login_url"`
	ShowCopyPastePrompt   bool                   `json:"show_copy_paste_prompt"`
	ShuffleQuestions      bool                   `json:"shuffle_questions"`
	Starred               bool                   `json:"starred"`
	Starttime             string                 `json:"starttime"`
	State                 string                 `json:"state"`
	Tags                  []string               `json:"tags"`
	TestAdmins            []string               `json:"test_admins"`
	TrackEditorPaste      bool                   `json:"track_editor_paste"`
	UniqueId              string                 `json:"unique_id"`
}
type TestsShow struct {
	CandidateDetails      []string               `json:"candidate_details"`
	CandidateTabSwitch    bool                   `json:"candidate_tab_switch"`
	CreatedAt             string                 `json:"created_at"`
	CustomAcknowledgeText string                 `json:"custom_acknowledge_text"`
	CutoffScore           float32                `json:"cutoff_score"`
	Draft                 bool                   `json:"draft"`
	Duration              float32                `json:"duration"`
	EnableAcknowledgement bool                   `json:"enable_acknowledgement"`
	EnableProctoring      bool                   `json:"enable_proctoring"`
	Endtime               string                 `json:"endtime"`
	Experience            []string               `json:"experience"`
	HideCompileTest       bool                   `json:"hide_compile_test"`
	HideTemplate          bool                   `json:"hide_template"`
	Id                    string                 `json:"id"`
	IdeConfig             string                 `json:"ide_config"`
	Instructions          string                 `json:"instructions"`
	Languages             []string               `json:"languages"`
	Locked                bool                   `json:"locked"`
	LockedBy              string                 `json:"locked_by"`
	MasterPassword        string                 `json:"master_password"`
	McqCorrectScore       float32                `json:"mcq_correct_score"`
	McqIncorrectScore     float32                `json:"mcq_incorrect_score"`
	Name                  string                 `json:"name"`
	Owner                 string                 `json:"owner"`
	PublicLoginUrl        string                 `json:"public_login_url"`
	Questions             []string               `json:"questions"`
	RoleIds               []string               `json:"role_ids"`
	Sections              map[string]interface{} `json:"sections"`
	Secure                bool                   `json:"secure"`
	ShortLoginUrl         string                 `json:"short_login_url"`
	ShowCopyPastePrompt   bool                   `json:"show_copy_paste_prompt"`
	ShuffleQuestions      bool                   `json:"shuffle_questions"`
	Starred               bool                   `json:"starred"`
	Starttime             string                 `json:"starttime"`
	State                 string                 `json:"state"`
	Tags                  []string               `json:"tags"`
	TestAdmins            []string               `json:"test_admins"`
	TrackEditorPaste      bool                   `json:"track_editor_paste"`
	UniqueId              string                 `json:"unique_id"`
}
type TestsUpdate struct {
	CandidateDetails      []string `json:"candidate_details"`
	CandidateTabSwitch    bool     `json:"candidate_tab_switch"`
	CustomAcknowledgeText string   `json:"custom_acknowledge_text"`
	CutoffScore           float32  `json:"cutoff_score"`
	Draft                 bool     `json:"draft"`
	Duration              float32  `json:"duration"`
	EnableAcknowledgement bool     `json:"enable_acknowledgement"`
	EnableProctoring      bool     `json:"enable_proctoring"`
	Endtime               string   `json:"endtime"`
	Experience            []string `json:"experience"`
	HideCompileTest       bool     `json:"hide_compile_test"`
	HideTemplate          bool     `json:"hide_template"`
	IdeConfig             string   `json:"ide_config"`
	Instructions          string   `json:"instructions"`
	Languages             []string `json:"languages"`
	Locked                bool     `json:"locked"`
	MasterPassword        string   `json:"master_password"`
	McqCorrectScore       float32  `json:"mcq_correct_score"`
	McqIncorrectScore     float32  `json:"mcq_incorrect_score"`
	Name                  string   `json:"name"`
	Questions             []string `json:"questions"`
	RoleIds               []string `json:"role_ids"`
	Secure                bool     `json:"secure"`
	ShowCopyPastePrompt   bool     `json:"show_copy_paste_prompt"`
	ShuffleQuestions      bool     `json:"shuffle_questions"`
	Starttime             string   `json:"starttime"`
	Tags                  []string `json:"tags"`
	TestAdmins            []string `json:"test_admins"`
	TrackEditorPaste      bool     `json:"track_editor_paste"`
}
type UserCreate struct {
	CandidatesPermission       float32          `json:"candidates_permission"`
	CompanyAdmin               bool             `json:"company_admin"`
	Country                    string           `json:"country"`
	Email                      string           `json:"email"`
	Firstname                  string           `json:"firstname"`
	InterviewsPermission       float32          `json:"interviews_permission"`
	Lastname                   string           `json:"lastname"`
	Phone                      string           `json:"phone"`
	QuestionsPermission        float32          `json:"questions_permission"`
	Role                       string           `json:"role"`
	SendEmail                  bool             `json:"send_email"`
	SharedCandidatesPermission float32          `json:"shared_candidates_permission"`
	SharedInterviewsPermission float32          `json:"shared_interviews_permission"`
	SharedQuestionsPermission  float32          `json:"shared_questions_permission"`
	SharedTestsPermission      float32          `json:"shared_tests_permission"`
	TeamAdmin                  bool             `json:"team_admin"`
	Teams                      []UserTeamCreate `json:"teams"`
	TestsPermission            float32          `json:"tests_permission"`
}
type UserIndex struct {
	Activated                  bool     `json:"activated"`
	CandidatesPermission       float32  `json:"candidates_permission"`
	CompanyAdmin               bool     `json:"company_admin"`
	Country                    string   `json:"country"`
	Email                      string   `json:"email"`
	Firstname                  string   `json:"firstname"`
	Id                         string   `json:"id"`
	InterviewsPermission       float32  `json:"interviews_permission"`
	LastActivityTime           string   `json:"last_activity_time"`
	Lastname                   string   `json:"lastname"`
	Phone                      string   `json:"phone"`
	QuestionsPermission        float32  `json:"questions_permission"`
	Role                       string   `json:"role"`
	SharedCandidatesPermission float32  `json:"shared_candidates_permission"`
	SharedInterviewsPermission float32  `json:"shared_interviews_permission"`
	SharedQuestionsPermission  float32  `json:"shared_questions_permission"`
	SharedTestsPermission      float32  `json:"shared_tests_permission"`
	Status                     string   `json:"status"`
	TeamAdmin                  bool     `json:"team_admin"`
	Teams                      []string `json:"teams"`
	TestsPermission            float32  `json:"tests_permission"`
	Timezone                   string   `json:"timezone"`
}
type UserRef struct {
	Activated                  bool    `json:"activated"`
	CandidatesPermission       float32 `json:"candidates_permission"`
	CompanyAdmin               bool    `json:"company_admin"`
	Country                    string  `json:"country"`
	Email                      string  `json:"email"`
	Firstname                  string  `json:"firstname"`
	Id                         string  `json:"id"`
	InterviewsPermission       float32 `json:"interviews_permission"`
	LastActivityTime           string  `json:"last_activity_time"`
	Lastname                   string  `json:"lastname"`
	Phone                      string  `json:"phone"`
	QuestionsPermission        float32 `json:"questions_permission"`
	Role                       string  `json:"role"`
	SendEmail                  bool    `json:"send_email"`
	SharedCandidatesPermission float32 `json:"shared_candidates_permission"`
	SharedInterviewsPermission float32 `json:"shared_interviews_permission"`
	SharedQuestionsPermission  float32 `json:"shared_questions_permission"`
	SharedTestsPermission      float32 `json:"shared_tests_permission"`
	Status                     string  `json:"status"`
	TeamAdmin                  bool    `json:"team_admin"`
	TestsPermission            float32 `json:"tests_permission"`
	Timezone                   string  `json:"timezone"`
}
type UserShow struct {
	Activated                  bool     `json:"activated"`
	CandidatesPermission       float32  `json:"candidates_permission"`
	CompanyAdmin               bool     `json:"company_admin"`
	Country                    string   `json:"country"`
	Email                      string   `json:"email"`
	Firstname                  string   `json:"firstname"`
	Id                         string   `json:"id"`
	InterviewsPermission       float32  `json:"interviews_permission"`
	LastActivityTime           string   `json:"last_activity_time"`
	Lastname                   string   `json:"lastname"`
	Phone                      string   `json:"phone"`
	QuestionsPermission        float32  `json:"questions_permission"`
	Role                       string   `json:"role"`
	SharedCandidatesPermission float32  `json:"shared_candidates_permission"`
	SharedInterviewsPermission float32  `json:"shared_interviews_permission"`
	SharedQuestionsPermission  float32  `json:"shared_questions_permission"`
	SharedTestsPermission      float32  `json:"shared_tests_permission"`
	Status                     string   `json:"status"`
	TeamAdmin                  bool     `json:"team_admin"`
	Teams                      []string `json:"teams"`
	TestsPermission            float32  `json:"tests_permission"`
	Timezone                   string   `json:"timezone"`
}
type UserTeamCreate struct {
	Id string `json:"id"`
}
type UserTeamMembershipCreate struct {
}
type UserTeamMembershipIndex struct {
	Team string `json:"team"`
	User string `json:"user"`
}
type UserTeamMembershipRef struct {
}
type UserTeamMembershipShow struct {
	Team string `json:"team"`
	User string `json:"user"`
}
type UserTeamMembershipUpdate struct {
}
type UserUpdate struct {
	CandidatesPermission       float32 `json:"candidates_permission"`
	CompanyAdmin               bool    `json:"company_admin"`
	Country                    string  `json:"country"`
	Firstname                  string  `json:"firstname"`
	InterviewsPermission       float32 `json:"interviews_permission"`
	Lastname                   string  `json:"lastname"`
	Phone                      string  `json:"phone"`
	QuestionsPermission        float32 `json:"questions_permission"`
	Role                       string  `json:"role"`
	SharedCandidatesPermission float32 `json:"shared_candidates_permission"`
	SharedInterviewsPermission float32 `json:"shared_interviews_permission"`
	SharedQuestionsPermission  float32 `json:"shared_questions_permission"`
	SharedTestsPermission      float32 `json:"shared_tests_permission"`
	TeamAdmin                  bool    `json:"team_admin"`
	TestsPermission            float32 `json:"tests_permission"`
}
type CreateGroupsParamsBody SCIMTeamCreate

type CreateGroupsParams struct {
	CreateGroupsParamsBody
}
type CreateGroupsResponse SCIMTeamShow

func (c *Client) CreateGroups(params *CreateGroupsParams) (*CreateGroupsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Groups")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response CreateGroupsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type DeleteGroupByIdParams struct {
	Id string
}
type DeleteGroupByIdResponse struct{}

func (c *Client) DeleteGroupById(params *DeleteGroupByIdParams) (*DeleteGroupByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Groups/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response DeleteGroupByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type GetGroupByIdParams struct {
	Id string
}
type GetGroupByIdResponse SCIMTeamShow

func (c *Client) GetGroupById(params *GetGroupByIdParams) (*GetGroupByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Groups/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response GetGroupByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type UpdateGroupByIdParamsBody SCIMPatchCreate

type UpdateGroupByIdParams struct {
	Id string
	UpdateGroupByIdParamsBody
}
type UpdateGroupByIdResponse struct {
	Message string        `json:"message"`
	Schemas []interface{} `json:"schemas"`
}

func (c *Client) UpdateGroupById(params *UpdateGroupByIdParams) (*UpdateGroupByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Groups/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PATCH", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response UpdateGroupByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type GetGroupsParams struct {
	Limit  int
	Offset int
}
type GetGroupsResponse struct {
	Resources    []SCIMTeamEntryIndex `json:"Resources"`
	ItemsPerPage float32              `json:"itemsPerPage"`
	Schemas      []interface{}        `json:"schemas"`
	StartIndex   float32              `json:"startIndex"`
	TotalResults float32              `json:"totalResults"`
}

func (c *Client) GetGroups(params *GetGroupsParams) (*GetGroupsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Groups?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response GetGroupsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type CreateUsersParamsBody SCIMUserCreate

type CreateUsersParams struct {
	CreateUsersParamsBody
}
type CreateUsersResponse struct{}

func (c *Client) CreateUsers(params *CreateUsersParams) (*CreateUsersResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Users")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response CreateUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type DeleteUserByIdParams struct {
	Id string
}
type DeleteUserByIdResponse struct{}

func (c *Client) DeleteUserById(params *DeleteUserByIdParams) (*DeleteUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response DeleteUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type GetUserByIdParams struct {
	Id SCIMUserShow
}
type GetUserByIdResponse SCIMUserShow

func (c *Client) GetUserById(params *GetUserByIdParams) (*GetUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response GetUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type UpdateUserByIdParamsBody SCIMPatchCreate

type UpdateUserByIdParams struct {
	Id string
	UpdateUserByIdParamsBody
}
type UpdateUserByIdResponse struct {
	Message string        `json:"message"`
	Schemas []interface{} `json:"schemas"`
}

func (c *Client) UpdateUserById(params *UpdateUserByIdParams) (*UpdateUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PATCH", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response UpdateUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type PutUserByIdParamsBody SCIMUserUpdate

type PutUserByIdParams struct {
	Id string
	PutUserByIdParamsBody
}
type PutUserByIdResponse SCIMUserShow

func (c *Client) PutUserById(params *PutUserByIdParams) (*PutUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response PutUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type GetUsersParams struct {
	Limit  int
	Offset int
}
type GetUsersResponse struct {
	Resources    []SCIMUserEntryIndex `json:"Resources"`
	ItemsPerPage float32              `json:"itemsPerPage"`
	Schemas      []interface{}        `json:"schemas"`
	StartIndex   float32              `json:"startIndex"`
	TotalResults float32              `json:"totalResults"`
}

func (c *Client) GetUsers(params *GetUsersParams) (*GetUsersResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/Users?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response GetUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateAtsCodepairParamsBody ATSCodePairCreate

type V3CreateAtsCodepairParams struct {
	V3CreateAtsCodepairParamsBody
}
type V3CreateAtsCodepairResponse InterviewShow

func (c *Client) V3CreateAtsCodepair(params *V3CreateAtsCodepairParams) (*V3CreateAtsCodepairResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/ats/codepair")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateAtsCodepairResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateAtsCodescreenParamsBody ATSCodeScreenCreate

type V3CreateAtsCodescreenParams struct {
	V3CreateAtsCodescreenParamsBody
}
type V3CreateAtsCodescreenResponse struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	TestLink string `json:"test_link"`
}

func (c *Client) V3CreateAtsCodescreen(params *V3CreateAtsCodescreenParams) (*V3CreateAtsCodescreenResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/ats/codescreen")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateAtsCodescreenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetAuditLogParams struct {
	Limit  int
	Offset int
}
type V3GetAuditLogResponse struct {
	Data      []AuditLogIndex `json:"data"`
	First     string          `json:"first"`
	Last      string          `json:"last"`
	Next      string          `json:"next"`
	Offset    float32         `json:"offset"`
	PageTotal float32         `json:"page_total"`
	Previous  string          `json:"previous"`
	Total     float32         `json:"total"`
}

func (c *Client) V3GetAuditLog(params *V3GetAuditLogParams) (*V3GetAuditLogResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/audit_log?limit={limit}&offset={offset}&user_id={user_id}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetAuditLogResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetInterviewsParams struct {
	Limit         int
	Offset        int
	CreatedAt     string `json:"created_at"`
	UpdatedAt     string `json:"updated_at"`
	EndedAt       string `json:"ended_at"`
	User          string `json:"user"`
	Interviewers  string `json:"interviewers"`
	Access        string `json:"access"`
	CurrentStatus string `json:"current_status"`
	OrderBy       string `json:"order_by"`
	OrderDir      string `json:"order_dir"`
}
type V3GetInterviewsResponse struct {
	Data      []InterviewIndex `json:"data"`
	First     string           `json:"first"`
	Last      string           `json:"last"`
	Next      string           `json:"next"`
	Offset    float32          `json:"offset"`
	PageTotal float32          `json:"page_total"`
	Previous  string           `json:"previous"`
	Total     float32          `json:"total"`
}

func (c *Client) V3GetInterviews(params *V3GetInterviewsParams) (*V3GetInterviewsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/interviews")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetInterviewsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateInterviewsParamsBody InterviewCreate

type V3CreateInterviewsParams struct {
	V3CreateInterviewsParamsBody
}
type V3CreateInterviewsResponse InterviewShow

func (c *Client) V3CreateInterviews(params *V3CreateInterviewsParams) (*V3CreateInterviewsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/interviews")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateInterviewsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteInterviewByInterviewIdParams struct {
	InterviewId string
}
type V3DeleteInterviewByInterviewIdResponse struct{}

func (c *Client) V3DeleteInterviewByInterviewId(params *V3DeleteInterviewByInterviewIdParams) (*V3DeleteInterviewByInterviewIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/interviews/{interview_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{interview_id}", fmt.Sprintf("%v", params.InterviewId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteInterviewByInterviewIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetInterviewByInterviewIdParams struct {
	InterviewId string
}
type V3GetInterviewByInterviewIdResponse InterviewShow

func (c *Client) V3GetInterviewByInterviewId(params *V3GetInterviewByInterviewIdParams) (*V3GetInterviewByInterviewIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/interviews/{interview_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{interview_id}", fmt.Sprintf("%v", params.InterviewId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetInterviewByInterviewIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutInterviewByInterviewIdParamsBody InterviewUpdate

type V3PutInterviewByInterviewIdParams struct {
	InterviewId string
	V3PutInterviewByInterviewIdParamsBody
}
type V3PutInterviewByInterviewIdResponse InterviewShow

func (c *Client) V3PutInterviewByInterviewId(params *V3PutInterviewByInterviewIdParams) (*V3PutInterviewByInterviewIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/interviews/{interview_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{interview_id}", fmt.Sprintf("%v", params.InterviewId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutInterviewByInterviewIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetQuestionsParams struct {
	Status     string `json:"status"`
	Access     string `json:"access"`
	Difficulty string `json:"difficulty"`
	Type       string `json:"type"`
	Owner      string `json:"owner"`
	Tags       string `json:"tags"`
	Skills     string `json:"skills"`
	Languages  string `json:"languages"`
}
type V3GetQuestionsResponse struct {
	Data      []QuestionIndex `json:"data"`
	First     string          `json:"first"`
	Last      string          `json:"last"`
	Next      string          `json:"next"`
	Offset    float32         `json:"offset"`
	PageTotal float32         `json:"page_total"`
	Previous  string          `json:"previous"`
	Total     float32         `json:"total"`
}

func (c *Client) V3GetQuestions(params *V3GetQuestionsParams) (*V3GetQuestionsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/questions")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetQuestionsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateQuestionsParamsBody QuestionCreate

type V3CreateQuestionsParams struct {
	V3CreateQuestionsParamsBody
}
type V3CreateQuestionsResponse QuestionShow

func (c *Client) V3CreateQuestions(params *V3CreateQuestionsParams) (*V3CreateQuestionsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/questions")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateQuestionsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetQuestionByQuestionIdParams struct {
	QuestionId string
}
type V3GetQuestionByQuestionIdResponse QuestionShow

func (c *Client) V3GetQuestionByQuestionId(params *V3GetQuestionByQuestionIdParams) (*V3GetQuestionByQuestionIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/questions/{question_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{question_id}", fmt.Sprintf("%v", params.QuestionId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetQuestionByQuestionIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutQuestionByQuestionIdParamsBody QuestionUpdate

type V3PutQuestionByQuestionIdParams struct {
	QuestionId string
	V3PutQuestionByQuestionIdParamsBody
}
type V3PutQuestionByQuestionIdResponse QuestionShow

func (c *Client) V3PutQuestionByQuestionId(params *V3PutQuestionByQuestionIdParams) (*V3PutQuestionByQuestionIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/questions/{question_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{question_id}", fmt.Sprintf("%v", params.QuestionId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutQuestionByQuestionIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutQuestionByQuestionIdGenerateParamsBody struct {
	AllowedLanguages string `json:"allowedLanguages"`
	FunctionName     string `json:"functionName"`
	FunctionParams   string `json:"functionParams"`
	FunctionReturn   string `json:"functionReturn"`
	Type             string `json:"type"`
}

type V3PutQuestionByQuestionIdGenerateParams struct {
	QuestionId string
	V3PutQuestionByQuestionIdGenerateParamsBody
}
type V3PutQuestionByQuestionIdGenerateResponse struct {
	AllowedLanguages    string `json:"allowedLanguages"`
	CTemplate           string `json:"c_template"`
	CTemplateHead       string `json:"c_template_head"`
	CTemplateTail       string `json:"c_template_tail"`
	ClojureTemplate     string `json:"clojure_template"`
	ClojureTemplateHead string `json:"clojure_template_head"`
	ClojureTemplateTail string `json:"clojure_template_tail"`
	FunctionName        string `json:"functionName"`
	FunctionParams      string `json:"functionParams"`
	FunctionReturn      string `json:"functionReturn"`
	TemplateType        string `json:"template_type"`
}

func (c *Client) V3PutQuestionByQuestionIdGenerate(params *V3PutQuestionByQuestionIdGenerateParams) (*V3PutQuestionByQuestionIdGenerateResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/questions/{question_id}/generate")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{question_id}", fmt.Sprintf("%v", params.QuestionId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutQuestionByQuestionIdGenerateResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateQuestionByQuestionIdTestcasesParamsBody struct {
	Explanation string  `json:"explanation"`
	Input       string  `json:"input"`
	Name        string  `json:"name"`
	Output      string  `json:"output"`
	Qid         float32 `json:"qid"`
	Sample      int     `json:"sample"`
	Score       float32 `json:"score"`
	Type        string  `json:"type"`
}

type V3CreateQuestionByQuestionIdTestcasesParams struct {
	QuestionId string
	V3CreateQuestionByQuestionIdTestcasesParamsBody
}
type V3CreateQuestionByQuestionIdTestcasesResponse struct{}

func (c *Client) V3CreateQuestionByQuestionIdTestcases(params *V3CreateQuestionByQuestionIdTestcasesParams) (*V3CreateQuestionByQuestionIdTestcasesResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/questions/{question_id}/testcases")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{question_id}", fmt.Sprintf("%v", params.QuestionId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateQuestionByQuestionIdTestcasesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateTeamsParamsBody TeamCreate

type V3CreateTeamsParams struct {
	V3CreateTeamsParamsBody
}
type V3CreateTeamsResponse TeamShow

func (c *Client) V3CreateTeams(params *V3CreateTeamsParams) (*V3CreateTeamsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateTeamsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteTeamByIdParams struct {
	Id string
}
type V3DeleteTeamByIdResponse struct{}

func (c *Client) V3DeleteTeamById(params *V3DeleteTeamByIdParams) (*V3DeleteTeamByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteTeamByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTeamByIdParams struct {
	Id string
}
type V3GetTeamByIdResponse TeamShow

func (c *Client) V3GetTeamById(params *V3GetTeamByIdParams) (*V3GetTeamByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTeamByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutTeamByIdParamsBody TeamUpdate

type V3PutTeamByIdParams struct {
	Id string
	V3PutTeamByIdParamsBody
}
type V3PutTeamByIdResponse struct{}

func (c *Client) V3PutTeamById(params *V3PutTeamByIdParams) (*V3PutTeamByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutTeamByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteTeamByTeamIdUserByUserIdParams struct {
	TeamId string
	UserId string
}
type V3DeleteTeamByTeamIdUserByUserIdResponse struct{}

func (c *Client) V3DeleteTeamByTeamIdUserByUserId(params *V3DeleteTeamByTeamIdUserByUserIdParams) (*V3DeleteTeamByTeamIdUserByUserIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{team_id}/users/{user_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{team_id}", fmt.Sprintf("%v", params.TeamId), -1)
	u.Path = strings.Replace(u.Path, "{user_id}", fmt.Sprintf("%v", params.UserId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteTeamByTeamIdUserByUserIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTeamByTeamIdUserByUserIdParams struct {
	TeamId string
	UserId string
}
type V3GetTeamByTeamIdUserByUserIdResponse UserTeamMembershipShow

func (c *Client) V3GetTeamByTeamIdUserByUserId(params *V3GetTeamByTeamIdUserByUserIdParams) (*V3GetTeamByTeamIdUserByUserIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{team_id}/users/{user_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{team_id}", fmt.Sprintf("%v", params.TeamId), -1)
	u.Path = strings.Replace(u.Path, "{user_id}", fmt.Sprintf("%v", params.UserId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTeamByTeamIdUserByUserIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateTeamByTeamIdUserByUserIdParams struct {
	TeamId  string
	UserId  string
	License string
}
type V3CreateTeamByTeamIdUserByUserIdResponse struct{}

func (c *Client) V3CreateTeamByTeamIdUserByUserId(params *V3CreateTeamByTeamIdUserByUserIdParams) (*V3CreateTeamByTeamIdUserByUserIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{team_id}/users/{user_id}?license={license}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{team_id}", fmt.Sprintf("%v", params.TeamId), -1)
	u.Path = strings.Replace(u.Path, "{user_id}", fmt.Sprintf("%v", params.UserId), -1)

	q := make(url.Values)
	q.Set("license", fmt.Sprintf("%v", params.License))

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateTeamByTeamIdUserByUserIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTeamByTeamIdUsersParams struct {
	TeamId string
	Limit  int
	Offset int
}
type V3GetTeamByTeamIdUsersResponse struct {
	Data      []UserTeamMembershipIndex `json:"data"`
	First     string                    `json:"first"`
	Last      string                    `json:"last"`
	Next      string                    `json:"next"`
	Offset    float32                   `json:"offset"`
	PageTotal float32                   `json:"page_total"`
	Previous  string                    `json:"previous"`
	Total     float32                   `json:"total"`
}

func (c *Client) V3GetTeamByTeamIdUsers(params *V3GetTeamByTeamIdUsersParams) (*V3GetTeamByTeamIdUsersResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams/{team_id}/users?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{team_id}", fmt.Sprintf("%v", params.TeamId), -1)

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTeamByTeamIdUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTeamsParams struct {
	Limit  int
	Offset int
}
type V3GetTeamsResponse struct {
	Data      []TeamIndex `json:"data"`
	First     string      `json:"first"`
	Last      string      `json:"last"`
	Next      string      `json:"next"`
	Offset    float32     `json:"offset"`
	PageTotal float32     `json:"page_total"`
	Previous  string      `json:"previous"`
	Total     float32     `json:"total"`
}

func (c *Client) V3GetTeams(params *V3GetTeamsParams) (*V3GetTeamsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/teams?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTeamsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTemplatesParams struct {
	Access string `json:"access"`
}
type V3GetTemplatesResponse struct {
	Data      []TemplateIndex `json:"data"`
	First     string          `json:"first"`
	Last      string          `json:"last"`
	Next      string          `json:"next"`
	Offset    float32         `json:"offset"`
	PageTotal float32         `json:"page_total"`
	Previous  string          `json:"previous"`
	Total     float32         `json:"total"`
}

func (c *Client) V3GetTemplates(params *V3GetTemplatesParams) (*V3GetTemplatesResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/templates")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTemplatesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTemplateByTemplateIdParams struct {
	TemplateId string
}
type V3GetTemplateByTemplateIdResponse TemplateShow

func (c *Client) V3GetTemplateByTemplateId(params *V3GetTemplateByTemplateIdParams) (*V3GetTemplateByTemplateIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/templates/{template_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{template_id}", fmt.Sprintf("%v", params.TemplateId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTemplateByTemplateIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateTestsParamsBody TestsCreate

type V3CreateTestsParams struct {
	V3CreateTestsParamsBody
}
type V3CreateTestsResponse TestsShow

func (c *Client) V3CreateTests(params *V3CreateTestsParams) (*V3CreateTestsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateTestsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteTestByIdParams struct {
	Id string
}
type V3DeleteTestByIdResponse struct{}

func (c *Client) V3DeleteTestById(params *V3DeleteTestByIdParams) (*V3DeleteTestByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteTestByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestByIdParams struct {
	Id string
}
type V3GetTestByIdResponse TestsShow

func (c *Client) V3GetTestById(params *V3GetTestByIdParams) (*V3GetTestByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutTestByIdParamsBody TestsUpdate

type V3PutTestByIdParams struct {
	Id string
	V3PutTestByIdParamsBody
}
type V3PutTestByIdResponse struct{}

func (c *Client) V3PutTestById(params *V3PutTestByIdParams) (*V3PutTestByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutTestByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateTestByIdArchiveParams struct {
	Id string
}
type V3CreateTestByIdArchiveResponse struct{}

func (c *Client) V3CreateTestByIdArchive(params *V3CreateTestByIdArchiveParams) (*V3CreateTestByIdArchiveResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{id}/archive")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateTestByIdArchiveResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestByIdInvitersParams struct {
	Id     string
	Limit  int
	Offset int
}
type V3GetTestByIdInvitersResponse struct {
	Data      []InvitersIndex `json:"data"`
	First     string          `json:"first"`
	Last      string          `json:"last"`
	Next      string          `json:"next"`
	Offset    float32         `json:"offset"`
	PageTotal float32         `json:"page_total"`
	Previous  string          `json:"previous"`
	Total     float32         `json:"total"`
}

func (c *Client) V3GetTestByIdInviters(params *V3GetTestByIdInvitersParams) (*V3GetTestByIdInvitersResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{id}/inviters")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestByIdInvitersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateTestByTestIdCandidatesParamsBody TestCandidateCreate

type V3CreateTestByTestIdCandidatesParams struct {
	TestId string
	V3CreateTestByTestIdCandidatesParamsBody
}
type V3CreateTestByTestIdCandidatesResponse struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	TestLink string `json:"test_link"`
}

func (c *Client) V3CreateTestByTestIdCandidates(params *V3CreateTestByTestIdCandidatesParams) (*V3CreateTestByTestIdCandidatesResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateTestByTestIdCandidatesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestByTestIdCandidatesSearchParams struct {
	TestId string
	Search string
	Limit  int
	Offset int
}
type V3GetTestByTestIdCandidatesSearchResponse struct {
	Data      []TestCandidateIndex `json:"data"`
	First     string               `json:"first"`
	Last      string               `json:"last"`
	Next      string               `json:"next"`
	Offset    float32              `json:"offset"`
	PageTotal float32              `json:"page_total"`
	Previous  string               `json:"previous"`
	Total     float32              `json:"total"`
}

func (c *Client) V3GetTestByTestIdCandidatesSearch(params *V3GetTestByTestIdCandidatesSearchParams) (*V3GetTestByTestIdCandidatesSearchResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates/search?search={search}&limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)

	q := make(url.Values)
	q.Set("search", fmt.Sprintf("%v", params.Search))
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestByTestIdCandidatesSearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutTestByTestIdCandidateByCandidateIdParamsBody TestCandidateUpdate

type V3PutTestByTestIdCandidateByCandidateIdParams struct {
	TestId      string
	CandidateId string
	V3PutTestByTestIdCandidateByCandidateIdParamsBody
}
type V3PutTestByTestIdCandidateByCandidateIdResponse TestCandidateShow

func (c *Client) V3PutTestByTestIdCandidateByCandidateId(params *V3PutTestByTestIdCandidateByCandidateIdParams) (*V3PutTestByTestIdCandidateByCandidateIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates/{candidate_id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)
	u.Path = strings.Replace(u.Path, "{candidate_id}", fmt.Sprintf("%v", params.CandidateId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutTestByTestIdCandidateByCandidateIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteTestByTestIdCandidateByCandidateIdInviteParams struct {
	TestId      string
	CandidateId string
}
type V3DeleteTestByTestIdCandidateByCandidateIdInviteResponse struct{}

func (c *Client) V3DeleteTestByTestIdCandidateByCandidateIdInvite(params *V3DeleteTestByTestIdCandidateByCandidateIdInviteParams) (*V3DeleteTestByTestIdCandidateByCandidateIdInviteResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates/{candidate_id}/invite")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)
	u.Path = strings.Replace(u.Path, "{candidate_id}", fmt.Sprintf("%v", params.CandidateId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteTestByTestIdCandidateByCandidateIdInviteResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestByTestIdCandidateByCandidateIdPdfParams struct {
	TestId      string
	CandidateId string
}
type V3GetTestByTestIdCandidateByCandidateIdPdfResponse struct{}

func (c *Client) V3GetTestByTestIdCandidateByCandidateIdPdf(params *V3GetTestByTestIdCandidateByCandidateIdPdfParams) (*V3GetTestByTestIdCandidateByCandidateIdPdfResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates/{candidate_id}/pdf?format=url")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)
	u.Path = strings.Replace(u.Path, "{candidate_id}", fmt.Sprintf("%v", params.CandidateId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestByTestIdCandidateByCandidateIdPdfResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteTestByTestIdCandidateByCandidateIdReportParams struct {
	TestId      string
	CandidateId string
}
type V3DeleteTestByTestIdCandidateByCandidateIdReportResponse struct{}

func (c *Client) V3DeleteTestByTestIdCandidateByCandidateIdReport(params *V3DeleteTestByTestIdCandidateByCandidateIdReportParams) (*V3DeleteTestByTestIdCandidateByCandidateIdReportResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates/{candidate_id}/report")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)
	u.Path = strings.Replace(u.Path, "{candidate_id}", fmt.Sprintf("%v", params.CandidateId), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteTestByTestIdCandidateByCandidateIdReportResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestByTestIdCandidateByCandidateIdParams struct {
	TestId           string
	CandidateId      string
	AdditionalFields string
}
type V3GetTestByTestIdCandidateByCandidateIdResponse TestCandidateShow

func (c *Client) V3GetTestByTestIdCandidateByCandidateId(params *V3GetTestByTestIdCandidateByCandidateIdParams) (*V3GetTestByTestIdCandidateByCandidateIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates/{candidate_id}?additional_fields={additional_fields}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)
	u.Path = strings.Replace(u.Path, "{candidate_id}", fmt.Sprintf("%v", params.CandidateId), -1)

	q := make(url.Values)
	q.Set("additional_fields", fmt.Sprintf("%v", params.AdditionalFields))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestByTestIdCandidateByCandidateIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestByTestIdCandidatesParams struct {
	TestId string
	Limit  int
	Offset int
}
type V3GetTestByTestIdCandidatesResponse struct {
	Data      []TestCandidateIndex `json:"data"`
	First     string               `json:"first"`
	Last      string               `json:"last"`
	Next      string               `json:"next"`
	Offset    float32              `json:"offset"`
	PageTotal float32              `json:"page_total"`
	Previous  string               `json:"previous"`
	Total     float32              `json:"total"`
}

func (c *Client) V3GetTestByTestIdCandidates(params *V3GetTestByTestIdCandidatesParams) (*V3GetTestByTestIdCandidatesResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests/{test_id}/candidates?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{test_id}", fmt.Sprintf("%v", params.TestId), -1)

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestByTestIdCandidatesResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetTestsParams struct {
	Limit  int
	Offset int
}
type V3GetTestsResponse struct {
	Data      []TestsIndex `json:"data"`
	First     string       `json:"first"`
	Last      string       `json:"last"`
	Next      string       `json:"next"`
	Offset    float32      `json:"offset"`
	PageTotal float32      `json:"page_total"`
	Previous  string       `json:"previous"`
	Total     float32      `json:"total"`
}

func (c *Client) V3GetTests(params *V3GetTestsParams) (*V3GetTestsResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/tests?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetTestsResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3CreateUsersParamsBody UserCreate

type V3CreateUsersParams struct {
	V3CreateUsersParamsBody
}
type V3CreateUsersResponse UserShow

func (c *Client) V3CreateUsers(params *V3CreateUsersParams) (*V3CreateUsersResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/users")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("POST", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3CreateUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetUsersSearchParams struct {
	SearchString string
	Limit        int
	Offset       int
}
type V3GetUsersSearchResponse struct {
	Data      []UserIndex `json:"data"`
	First     string      `json:"first"`
	Last      string      `json:"last"`
	Next      string      `json:"next"`
	Offset    float32     `json:"offset"`
	PageTotal float32     `json:"page_total"`
	Previous  string      `json:"previous"`
	Total     float32     `json:"total"`
}

func (c *Client) V3GetUsersSearch(params *V3GetUsersSearchParams) (*V3GetUsersSearchResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/users/search?search={search_string}&limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("search_string", fmt.Sprintf("%v", params.SearchString))
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetUsersSearchResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3DeleteUserByIdParams struct {
	Id string
}
type V3DeleteUserByIdResponse struct{}

func (c *Client) V3DeleteUserById(params *V3DeleteUserByIdParams) (*V3DeleteUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("DELETE", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3DeleteUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetUserByIdParams struct {
	Id UserShow
}
type V3GetUserByIdResponse UserShow

func (c *Client) V3GetUserById(params *V3GetUserByIdParams) (*V3GetUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3PutUserByIdParamsBody UserUpdate

type V3PutUserByIdParams struct {
	Id string
	V3PutUserByIdParamsBody
}
type V3PutUserByIdResponse struct{}

func (c *Client) V3PutUserById(params *V3PutUserByIdParams) (*V3PutUserByIdResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/users/{id}")
	if err != nil {
		return nil, err
	}
	u.Path = strings.Replace(u.Path, "{id}", fmt.Sprintf("%v", params.Id), -1)

	q := make(url.Values)

	u.RawQuery = q.Encode()
	requestBody, err := json.Marshal(params)
	if err != nil {
		return nil, err
	}
	buffer := bytes.NewBuffer(requestBody)

	req, err := http.NewRequest("PUT", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3PutUserByIdResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

type V3GetUsersParams struct {
	Limit  int
	Offset int
}
type V3GetUsersResponse struct {
	Data      []UserIndex `json:"data"`
	First     string      `json:"first"`
	Last      string      `json:"last"`
	Next      string      `json:"next"`
	Offset    float32     `json:"offset"`
	PageTotal float32     `json:"page_total"`
	Previous  string      `json:"previous"`
	Total     float32     `json:"total"`
}

func (c *Client) V3GetUsers(params *V3GetUsersParams) (*V3GetUsersResponse, error) {
	u, err := url.Parse("https://www.hackerrank.com/x/api/v3/users?limit={limit}&offset={offset}")
	if err != nil {
		return nil, err
	}

	q := make(url.Values)
	q.Set("limit", fmt.Sprintf("%v", params.Limit))
	q.Set("offset", fmt.Sprintf("%v", params.Offset))

	u.RawQuery = q.Encode()
	buffer := bytes.NewBuffer([]byte{})

	req, err := http.NewRequest("GET", u.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", "application/json")

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response V3GetUsersResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
