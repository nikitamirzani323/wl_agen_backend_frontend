<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = "";
	export let table_body_font = "";
	export let token = "";
	export let listHome = [];
	export let listCurr = [];
	export let listBank = [];
	export let totalrecord = 0;
    let dispatch = createEventDispatcher();
	let title_page = "MASTER"
	let title_page_admin = "ADMIN"
	let title_page_agen = ""
    let sData = "";
    let sDataAdmin = "";
    let sDataAgen = "";
    let sDataAgenAdmin = "";
    let myModal_newentry = "";
    let flag_btnsave = true;
    let idcurr_field = "";
    let name_field = "";
    let owner_field = "";
    let email_field = "";
    let phone1_field = "";
    let phone2_field = "";
    let note_field = "";
    let bank_id_field = "";
    let bank_name_field = "";
    let bank_norek_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";


    let admin_idmaster_field = "";
    let admin_tipe_field = "";
    let admin_username_field = "";
    let admin_password_field = "";
    let admin_name_field = "";
    let admin_phone1_field = "";
    let admin_phone2_field = "";
    let admin_status_field = "";
    let admin_create_field = "";
    let admin_update_field = "";

    let agen_idmaster_field = "";
    let agen_idcurr_field = "";
    let agen_name_field = "";
    let agen_owner_field = "";
    let agen_email_field = "";
    let agen_phone1_field = "";
    let agen_phone2_field = "";
    let agen_note_field = "";
    let agen_bank_id_field = "";
    let agen_bank_name_field = "";
    let agen_bank_norek_field = "";
    let agen_status_field = "";
    let agen_create_field = "";
    let agen_update_field = "";

    //AGEN ADMIN
    let listAgenAdmin = "";
    let filterAgenAdmin = "";
    let searchAgenAdmin = "";
    let idmasteragen = "";
    let agenadmin_idagen_field = "";
    let agenadmin_tipe_field = "";
    let agenadmin_username_field = "";
    let agenadmin_password_field = "";
    let agenadmin_name_field = "";
    let agenadmin_phone1_field = "";
    let agenadmin_phone2_field = "";
    let agenadmin_status_field = "";
    let agenadmin_create_field = "";
    let agenadmin_update_field = "";


    let idrecord = "";
    let idrecordmasteradmin = "";
    let idrecordmasteragen = "";
    let idrecordmasteragenadmin = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
        if (searchAgenAdmin) {
            filterAgenAdmin = listAgenAdmin.filter(
                (item) =>
                    item.masteragenadmin_username
                        .toLowerCase()
                        .includes(searchAgenAdmin.toLowerCase())
            );
        } else {
            filterAgenAdmin = [...listAgenAdmin];
        }
    }
    
    const NewData = (e,id,idcurr,name,owner,email,phone1,phone2,note,status,idbank,norekbank,nmownerbank,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            idrecord = id
            idcurr_field = idcurr;
            name_field = name;
            owner_field = owner;
            email_field = email;
            phone1_field = phone1;
            phone2_field = phone2;
            note_field = note;
            bank_id_field = idbank;
            bank_name_field = nmownerbank;
            bank_norek_field = norekbank;
            status_field = status;
            create_field = create;
            update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const NewDataAdmin = (e,id,idmaster,tipe,username,name,phone1,phone2,status,create,update) => {
        sDataAdmin = e
        if(sDataAdmin == "New"){
            clearField_masteradmin();
            admin_idmaster_field = idmaster
        }else{
            sDataAdmin = "Edit"
            let temp = username.split("-");
            idrecordmasteradmin = id
            admin_idmaster_field = temp[0]
            admin_tipe_field = tipe;
            admin_username_field = temp[1];
            admin_name_field = name;
            admin_phone1_field = phone1;
            admin_phone2_field = phone2;
            admin_status_field = status;
            admin_create_field = create;
            admin_update_field = update;
            
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud_admin"));
        myModal_newentry.show();
        
    };
    const NewDataAgen = (e,id,idmaster,idcurr,name,owner,email,phone1,phone2,note,status,idbank,norekbank,nmownerbank,create,update) => {
        sDataAgen = e
        title_page_agen = "MASTER - " + idmaster
        if(sDataAgen == "New"){
            clearField_masteragen();
            agen_idmaster_field = idmaster
        }else{
            sDataAgen = "Edit"
            idrecordmasteragen = id
            agen_idmaster_field = idmaster;
            agen_idcurr_field = idcurr;
            agen_name_field = name;
            agen_owner_field = owner;
            agen_email_field = email;
            agen_phone1_field = phone1;
            agen_phone2_field = phone2;
            agen_note_field = note;
            agen_bank_id_field = idbank;
            agen_bank_name_field = nmownerbank;
            agen_bank_norek_field = norekbank;
            agen_status_field = status;
            agen_create_field = create;
            agen_update_field = update;
            
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud_agen"));
        myModal_newentry.show();
        
    };
    const ShowFormAgenAdmin = (e, id, tipe, username,name,phone1,phone2,status,create,update) => {
        sDataAgenAdmin = e;
        if (e == "Edit") {
            idrecordmasteragenadmin = id;
            agenadmin_tipe_field = tipe;
            agenadmin_username_field = username;
            agenadmin_name_field = name;
            agenadmin_phone1_field = phone1;
            agenadmin_phone2_field = phone2;
            agenadmin_status_field = status;
            agenadmin_create_field = create;
            agenadmin_update_field = update;
        } else {
            clearField_masteragenadmin();
        }

        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud_agenadmin"));
        myModal_newentry.show();
    };
    const showListAgenAdmin = (idagen) => {
        if(idagen != "" || idagen != undefined){
            idmasteragen = idagen;
            agenadmin_idagen_field = idagen;
            myModal_newentry = new bootstrap.Modal(document.getElementById("modalagenadmin"));
            myModal_newentry.show();
            call_agenadmin(idagen);
        }
        
    };
    async function call_agenadmin(idagen) {
        listAgenAdmin = [];
        const res = await fetch("/api/masteragenadmin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "MASTERAGEN-ADMIN",
                masteragen_idagen: idagen,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listAgenAdmin = [
                        ...listAgenAdmin,
                        {
                        masteragenadmin_no: no,
                        masteragenadmin_id: record[i]["masteragenadmin_id"],
                        masteragenadmin_tipe: record[i]["masteragenadmin_tipe"],
                        masteragenadmin_username: record[i]["masteragenadmin_username"],
                        masteragenadmin_lastlogin: record[i]["masteragenadmin_lastlogin"],
                        masteragenadmin_name: record[i]["masteragenadmin_name"],
                        masteragenadmin_phone1: record[i]["masteragenadmin_phone1"],
                        masteragenadmin_phone2: record[i]["masteragenadmin_phone2"],
                        masteragenadmin_status: record[i]["masteragenadmin_status"],
                        masteragenadmin_status_css: record[i]["masteragenadmin_status_css"],
                        masteragenadmin_create: record[i]["masteragenadmin_create"],
                        masteragenadmin_update: record[i]["masteragenadmin_update"],
                        },
                    ];
                }
            }
        }
    }
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    const GenerateString = (y,e) => {
        idrecord = genRandomString(y,e)
    };
    const GenerateNumber = (y,e) => {
        admin_username_field = genRandomString(y,e)
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrecord.length < 4){
                flag = false
                msg += "The ID is maxlength 4\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Default Curreny is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
            if(bank_id_field == ""){
                flag = false
                msg += "The Bank is required\n"
            }
            if(bank_name_field == ""){
                flag = false
                msg += "The Bank Owner is required\n"
            }
            if(bank_norek_field == ""){
                flag = false
                msg += "The Bank Number is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Default Curreny is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
            if(bank_id_field == ""){
                flag = false
                msg += "The Bank is required\n"
            }
            if(bank_name_field == ""){
                flag = false
                msg += "The Bank Owner is required\n"
            }
            if(bank_norek_field == ""){
                flag = false
                msg += "The Bank Number is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/mastersave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"MASTER-SAVE",
                    master_id: idrecord,
                    master_idcurr: idcurr_field,
                    master_name: name_field,
                    master_owner: owner_field,
                    master_phone1: phone1_field,
                    master_phone2: phone2_field,
                    master_email: email_field,
                    master_note: note_field,
                    master_bank_id: bank_id_field,
                    master_bank_norek: bank_name_field,
                    master_bank_name: bank_norek_field,
                    master_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleMasterAdminSave() {
        let flag = true
        let msg = ""
        if(sDataAdmin == "New"){
            if(admin_idmaster_field == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(admin_username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(admin_password_field == ""){
                flag = false
                msg += "The Password is required\n"
            }
            if(admin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(admin_phone1_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(admin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecordmasteradmin == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(admin_idmaster_field == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(admin_username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            
            if(admin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(admin_phone1_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(admin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/masteradminsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataAdmin,
                    page:"MASTER-SAVE",
                    masteradmin_id: parseInt(idrecordmasteradmin),
                    masteradmin_idmaster: admin_idmaster_field,
                    masteradmin_tipe: admin_tipe_field,
                    masteradmin_username: admin_idmaster_field+"-"+admin_username_field,
                    masteradmin_password: admin_password_field,
                    masteradmin_name: admin_name_field,
                    masteradmin_phone1: admin_phone1_field,
                    masteradmin_phone2: admin_phone2_field,
                    masteradmin_status: admin_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataAdmin=="New"){
                    clearField_masteradmin()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleAgenSave() {
        let flag = true
        let msg = ""
        if(sDataAgen == "New"){
            if(agen_idmaster_field == ""){
                flag = false
                msg += "The Master is required\n"
            }
            if(agen_idcurr_field == ""){
                flag = false
                msg += "The Default Curreny is required\n"
            }
            if(agen_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(agen_owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(agen_phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(agen_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
            if(agen_bank_id_field == ""){
                flag = false
                msg += "The Bank is required\n"
            }
            if(agen_bank_name_field == ""){
                flag = false
                msg += "The Bank Owner is required\n"
            }
            if(agen_bank_norek_field == ""){
                flag = false
                msg += "The Bank Number is required\n"
            }
        }else{
            if(idrecordmasteragen == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(agen_idmaster_field == ""){
                flag = false
                msg += "The Master is required\n"
            }
            if(agen_idcurr_field == ""){
                flag = false
                msg += "The Default Curreny is required\n"
            }
            if(agen_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(agen_owner_field == ""){
                flag = false
                msg += "The Owner is required\n"
            }
            if(agen_phone1_field == ""){
                flag = false
                msg += "The Phone 1 is required\n"
            }
            if(agen_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
            if(agen_bank_id_field == ""){
                flag = false
                msg += "The Bank is required\n"
            }
            if(agen_bank_name_field == ""){
                flag = false
                msg += "The Bank Owner is required\n"
            }
            if(agen_bank_norek_field == ""){
                flag = false
                msg += "The Bank Number is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/masteragensave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataAgen,
                    page:"MASTER-SAVE",
                    masteragen_id: idrecordmasteragen,
                    masteragen_idmaster: agen_idmaster_field,
                    masteragen_idcurr: agen_idcurr_field,
                    masteragen_name: agen_name_field,
                    masteragen_owner: agen_owner_field,
                    masteragen_phone1: agen_phone1_field,
                    masteragen_phone2: agen_phone2_field,
                    masteragen_email: agen_email_field,
                    masteragen_note: agen_note_field,
                    masteragen_bank_id: agen_bank_id_field,
                    masteragen_bank_norek: agen_bank_norek_field,
                    masteragen_bank_name: agen_bank_name_field,
                    masteragen_status: agen_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataAgen=="New"){
                    clearField_masteragen()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleMasterAgenAdminSave() {
        let flag = true
        let msg = ""
        if(sDataAgenAdmin == "New"){
            if(agenadmin_idagen_field == ""){
                flag = false
                msg += "The Code Agen is required\n"
            }
            if(agenadmin_tipe_field == ""){
                flag = false
                msg += "The Tipe is required\n"
            }
            if(agenadmin_username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(agenadmin_password_field == ""){
                flag = false
                msg += "The Password is required\n"
            }
            if(agenadmin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(agenadmin_phone1_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(agenadmin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecordmasteragenadmin == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(agenadmin_idagen_field == ""){
                flag = false
                msg += "The Code Agen is required\n"
            }
            if(agenadmin_tipe_field == ""){
                flag = false
                msg += "The Tipe is required\n"
            }
            if(agenadmin_username_field == ""){
                flag = false
                msg += "The Username is required\n"
            }
            if(agenadmin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(agenadmin_phone1_field == ""){
                flag = false
                msg += "The Phone is required\n"
            }
            if(agenadmin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/masteragenadminsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataAgenAdmin,
                    page:"MASTER-SAVE",
                    masteragenadmin_id: idrecordmasteragenadmin,
                    masteragenadmin_idmasteragen: agenadmin_idagen_field,
                    masteragenadmin_tipe: agenadmin_tipe_field,
                    masteragenadmin_username: agenadmin_username_field,
                    masteragenadmin_password: agenadmin_password_field,
                    masteragenadmin_name: agenadmin_name_field,
                    masteragenadmin_phone1: agenadmin_phone1_field,
                    masteragenadmin_phone2: agenadmin_phone2_field,
                    masteragenadmin_status: agenadmin_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataAgenAdmin=="New"){
                    clearField_masteragenadmin()
                }
                msgloader = json.message;
                call_agenadmin(idmasteragen)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    function clearField(){
        idrecord = "";
        idcurr_field = "";
        name_field = "";
        owner_field = "";
        email_field = "";
        phone1_field = "";
        phone2_field = "";
        note_field = "";
        status_field = "";
        bank_id_field = "";
        bank_name_field = "";
        bank_norek_field = "";
        create_field = "";
        update_field = "";
    }
    function clearField_masteradmin(){
        admin_tipe_field = "";
        admin_username_field = "";
        admin_password_field = "";
        admin_name_field = "";
        admin_phone1_field = "";
        admin_phone2_field = "";
        admin_status_field = "";
        admin_create_field = "";
        admin_update_field = "";
    }
    function clearField_masteragen(){
        agen_idmaster_field = "";
        agen_idcurr_field = "";
        agen_name_field = "";
        agen_owner_field = "";
        agen_email_field = "";
        agen_phone1_field = "";
        agen_phone2_field = "";
        agen_note_field = "";
        agen_bank_id_field = "";
        agen_bank_name_field = "";
        agen_bank_norek_field = "";
        agen_status_field = "";
        agen_create_field = "";
        agen_update_field = "";
    }
    function clearField_masteragenadmin(){
        idrecordmasteragenadmin = "";
        agenadmin_tipe_field = "";
        agenadmin_username_field = "";
        agenadmin_password_field = "";
        agenadmin_name_field = "";
        agenadmin_phone1_field = "";
        agenadmin_phone2_field = "";
        agenadmin_status_field = "";
        agenadmin_create_field = "";
        agenadmin_update_field = "";
    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","","","","","","","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
            case "FORMNEW_AGENADMIN":
                ShowFormAgenAdmin("New");
                break;
        }
    }
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
  
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }
    function uperCase(element) {
        function onInput(event) {
            element.value = element.value.toUpperCase();
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
    function lowerCase(element) {
        function onInput(event) {
            element.value = element.value.toLowerCase();
        }
        element.addEventListener("input", onInput);
        return {
            destroy() {
                element.removeEventListener("input", onInput);
            },
        };
    }
    function genRandomString(tipe,length) {
        let chars = '';
        switch(tipe){
            case "UPPERCASE":
                chars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';break;
            case "LOWERCASE":
                chars = 'abcdefghijklmnopqrstuvwxyz';break;
            case "NUMBER":
                chars = '0123456789';break;
        }
        let charLength = chars.length;
        let result = '';
        for ( let i = 0; i < length; i++ ) {
            result += chars.charAt(Math.floor(Math.random() * charLength));
        }
        return result;
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container-fluid" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button on:click={callFunction}
                button_function="NEW"
                button_title="New"
                button_css="btn-dark"/>
            <Button on:click={callFunction}
                button_function="REFRESH"
                button_title="Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Master"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="3">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
                                <th NOWRAP width="7%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">START</th>
                                <th NOWRAP width="7%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">END</th>
                                <th NOWRAP width="3%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CODE</th>
                                <th NOWRAP width="3%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CURR</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MASTER</th>
                                <th NOWRAP width="35%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">AGEN</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewData("Edit",rec.home_id, rec.home_idcurr,
                                                rec.home_name,rec.home_owner,rec.home_email,rec.home_phone1,rec.home_phone2,rec.home_note,rec.home_status,
                                                rec.home_bank_id,rec.home_bank_norek,rec.home_bank_name,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                NewDataAdmin("New",idrecordmasteradmin,rec.home_id);
                                            }} class="bi bi-person-plus"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                                //e,id,idmaster,idcurr,name,owner,email,phone1,phone2,note,status,idbank,norekbank,nmownerbank,create,update
                                                NewDataAgen("New",idrecordmasteragen,rec.home_id,rec.home_idcurr);
                                            }} class="bi bi-database-add"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_start}</td>
                                    <td nowrap style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_end}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_idcurr}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <b>{rec.home_name}</b><br />
                                        OWNER : {rec.home_owner} <br />
                                        EMAIL : {rec.home_email} <br />
                                        PHONE : {rec.home_phone1} / {rec.home_phone2}<br />
                                        {rec.home_bank_id} - {rec.home_bank_norek} - {rec.home_bank_name}
                                        <br />
                                        <table class="table table-bordered">
                                            <thead>
                                                <tr>
                                                    <th style="text-align: center;vertical-align: top;font-size: 11px;" colspan=4>ADMIN</th>
                                                </tr>
                                                <tr>
                                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-size: 11px;">STATUS</th>
                                                    <th NOWRAP width="3%" style="text-align: left;vertical-align: top;font-size: 11px;">TIPE</th>
                                                    <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-size: 11px;">USERNAME</th>
                                                    <th width="*" style="text-align: left;vertical-align: top;font-size: 11px;">NAME</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {#each rec.home_listadmin as rec2}
                                                <tr>
                                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec2.masteradmin_status_css}">
                                                            {status(rec2.masteradmin_status)}
                                                        </span>
                                                    </td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.masteradmin_tipe}</td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">
                                                        <span
                                                            on:click={() => {
                                                                //e,id,idmaster,tipe,username,name,phone1,phone2,status,create,update
                                                                NewDataAdmin("Edit",rec2.masteradmin_id,rec2.masteradmin_idmaster,
                                                                rec2.masteradmin_tipe,rec2.masteradmin_username,rec2.masteradmin_name,
                                                                rec2.masteradmin_phone1,rec2.masteradmin_phone2,rec2.masteradmin_status,
                                                                rec2.masteradmin_create,rec2.masteradmin_update);
                                                            }} 
                                                            style="cursor: pointer;text-decoration:underline;">{rec2.masteradmin_username}</span>
                                                    </td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.masteradmin_name}</td>
                                                </tr>
                                                {/each}
                                            </tbody>
                                        </table>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                        <table class="table table-bordered">
                                            <thead>
                                                <tr>
                                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="1">&nbsp;</th>
                                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-size: 11px;">STATUS</th>
                                                    <th width="3%" style="text-align: left;vertical-align: top;font-size: 11px;">CODE</th>
                                                    <th width="*" style="text-align: left;vertical-align: top;font-size: 11px;">NAME</th>
                                                </tr>
                                            </thead>
                                            <tbody>
                                                {#each rec.home_listagen as rec2}
                                                <tr>
                                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                                        <i on:click={() => {
                                                                showListAgenAdmin(rec2.masteragen_id);
                                                            }} class="bi bi-person-plus"></i>
                                                    </td>
                                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec2.masteragen_status_css}">
                                                            {status(rec2.masteragen_status)}
                                                        </span>
                                                    </td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">{rec2.masteragen_id}</td>
                                                    <td NOWRAP  style="text-align: left;vertical-align: top;font-size: 11px;">
                                                        <span
                                                            on:click={() => {
                                                                //e,id,idmaster,idcurr,name,owner,email,phone1,phone2,note,status,idbank,norekbank,nmownerbank,create,update
                                                                NewDataAgen("Edit",rec2.masteragen_id,rec.home_id,
                                                                rec2.masteragen_idcurr,rec2.masteragen_nmagen,rec2.masteragen_owner,
                                                                rec2.masteragen_email,rec2.masteragen_phone1,rec2.masteragen_phone2,
                                                                rec2.masteragen_note,rec2.masteragen_status,rec2.masteragen_bank_id,
                                                                rec2.masteragen_bank_norek,rec2.masteragen_bank_name,
                                                                rec2.masteragen_create,rec2.masteragen_update);
                                                            }} 
                                                            style="cursor: pointer;text-decoration:underline;">{rec2.masteragen_nmagen}</span>
                                                    </td>
                                                </tr>
                                                {/each}
                                            </tbody>
                                        </table>
                                    </td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    <div class="input-group mb-3">
                        <input bind:value={idrecord}
                            use:uperCase
                            class="required form-control"
                            maxlength="3"
                            disabled
                            type="text"
                            placeholder="CODE"/>
                        {#if sData != "Edit"}
                        <button on:click={() => {
                                GenerateString("UPPERCASE",4);
                            }}
                            type="button" class="btn btn-info">Generate</button>
                        {/if}
                      </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Default Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        name="currency" id="currency" 
                        class="required form-control">
                        {#each listCurr as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input bind:value={owner_field}
                        class="required"
                        type="text"
                        placeholder="Owner"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input bind:value={email_field}
                        class=""
                        type="text"
                        placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <Input bind:value={phone1_field}
                        class="required"
                        type="text"
                        placeholder="Phone 1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <Input bind:value={phone2_field}
                        class=""
                        type="text"
                        placeholder="Phone 2"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Bank</label>
                    <select
                        bind:value="{bank_id_field}" 
                        name="bankid" id="bankid" 
                        class="required form-control">
                        {#each listBank as rec}
                        <option value="{rec.bank_id}">{rec.bank_category} - {rec.bank_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Nomor Rekening</label>
                    <Input bind:value={bank_norek_field}
                        class="required"
                        type="text"
                        placeholder="Bank Nomor Rekening"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Nama Bank</label>
                    <Input bind:value={bank_name_field}
                        class="required"
                        type="text"
                        placeholder="Bank Nama Pemilik Rekening"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea style="height: 100px;resize: none;" bind:value={note_field} class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {create_field}<br />
                            Update : {update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalentrycrud_admin"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="MASTER {title_page_admin+"/"+sDataAdmin}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Tipe</label>
                    <select
                        bind:value="{admin_tipe_field}" 
                        name="currency" id="Tipe" 
                        class="required form-control">
                        <option value="MASTER">MASTER</option>
                        <option value="ADMIN">ADMIN</option>
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">username</label>
                    <div class="input-group mb-3">
                        <span class="input-group-text" id="basic-addon1">{admin_idmaster_field}</span>
                        <input bind:value={admin_username_field}
                            use:uperCase
                            class="required form-control"
                            maxlength="5"
                            disabled
                            type="text"
                            placeholder="CODE"/>
                        {#if sDataAdmin != "Edit"}
                        <button on:click={() => {
                                GenerateNumber("NUMBER",5);
                            }}
                            type="button" class="btn btn-info">Generate</button>
                        {/if}
                      </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Password</label>
                    <input
                        bind:value={admin_password_field}
                        type="password"
                        class="form-control "
                        placeholder="Password"
                        aria-label="Password"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={admin_name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <Input bind:value={admin_phone1_field}
                        class="required"
                        type="text"
                        placeholder="Phone 1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <Input bind:value={admin_phone2_field}
                        class=""
                        type="text"
                        placeholder="Phone 2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={admin_status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {admin_create_field}<br />
                            Update : {admin_update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleMasterAdminSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalentrycrud_agen"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page_agen+"/"+sDataAgen}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Default Currency</label>
                    <select
                        bind:value="{agen_idcurr_field}" 
                        name="currency" id="currency" 
                        class="required form-control">
                        {#each listCurr as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={agen_name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Owner</label>
                    <Input bind:value={agen_owner_field}
                        class="required"
                        type="text"
                        placeholder="Owner"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Email</label>
                    <Input bind:value={agen_email_field}
                        class=""
                        type="text"
                        placeholder="Email"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <Input bind:value={agen_phone1_field}
                        class="required"
                        type="text"
                        placeholder="Phone 1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <Input bind:value={agen_phone2_field}
                        class=""
                        type="text"
                        placeholder="Phone 2"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Bank</label>
                    <select
                        bind:value="{agen_bank_id_field}" 
                        name="bankid" id="bankid" 
                        class="required form-control">
                        {#each listBank as rec}
                        <option value="{rec.bank_id}">{rec.bank_category} - {rec.bank_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Nomor Rekening</label>
                    <Input bind:value={agen_bank_norek_field}
                        class="required"
                        type="text"
                        placeholder="Bank Nomor Rekening"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Nama Bank</label>
                    <Input bind:value={agen_bank_name_field}
                        class="required"
                        type="text"
                        placeholder="Bank Nama Pemilik Rekening"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Note</label>
                    <textarea style="height: 100px;resize: none;" bind:value={agen_note_field} class="form-control"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={agen_status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {agen_create_field}<br />
                            Update : {agen_update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleAgenSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
  modal_id="modalagenadmin"
  modal_size="modal-dialog-centered modal-lg"
  modal_title="ADMIN - {idmasteragen}"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={true}
  modal_search={true}>
  <slot:template slot="search">
    <div class="col-lg-12" style="padding: 5px;">
      <input
        bind:value={searchAgenAdmin}
        type="text"
        class="form-control"
        placeholder="Search Username"
        aria-label="Search"/>
    </div>
  </slot:template>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
            <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STATUS</th>
            <th width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TYPE</th>
            <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">USERNAME</th>
            <th width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">LASTLOGIN</th>
        </tr>
      </thead>
      <tbody>
        {#each filterAgenAdmin as rec}
          <tr on:click={() => {
                //e, id, tipe, username,name,phone1,phone2,status,create,update
              ShowFormAgenAdmin("Edit",rec.masteragenadmin_id, rec.masteragenadmin_tipe,rec.masteragenadmin_username,
              rec.masteragenadmin_name,rec.masteragenadmin_phone1,rec.masteragenadmin_phone2,rec.masteragenadmin_status,
              rec.masteragenadmin_create,rec.masteragenadmin_update);
            }} style="color:blue;text-decoration:underline;cursor:pointer;">
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.masteragenadmin_status_css}">
                    {status(rec.masteragenadmin_status)}
                </span>
            </td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.masteragenadmin_tipe}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.masteragenadmin_username}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.masteragenadmin_lastlogin}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
  <slot:template slot="footer">
    <Button
      on:click={callFunction}
      button_function="FORMNEW_AGENADMIN"
      button_title="New"
      button_css="btn-primary"/>
  </slot:template>
</Modal>

<Modal
	modal_id="modalentrycrud_agenadmin"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="AGEN - {idmasteragen} - {title_page_admin+"/"+sDataAgenAdmin}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Tipe</label>
                    <select
                        bind:value="{agenadmin_tipe_field}" 
                        name="currency" id="Tipe" 
                        class="required form-control">
                        <option value="MASTER">MASTER</option>
                        <option value="ADMIN">ADMIN</option>
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">username</label>
                    {#if sDataAgenAdmin == "New"}
                    <input bind:value={agenadmin_username_field}
                        use:lowerCase
                        class="required form-control"
                        maxlength="20"
                        type="text"
                        placeholder="USERNAME"/>
                    {:else}
                    <input bind:value={agenadmin_username_field}
                        use:lowerCase
                        disabled
                        class="required form-control"
                        maxlength="20"
                        type="text"
                        placeholder="USERNAME"/>
                    {/if}
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Password</label>
                    <input
                        bind:value={agenadmin_password_field}
                        type="password"
                        maxlength="30"
                        class="form-control "
                        placeholder="Password"
                        aria-label="Password"/>
                </div>
            </div>
            <div class="col-sm-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input bind:value={agenadmin_name_field}
                        class="required"
                        type="text"
                        placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 1</label>
                    <Input bind:value={agenadmin_phone1_field}
                        class="required"
                        type="text"
                        placeholder="Phone 1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone 2</label>
                    <Input bind:value={agenadmin_phone2_field}
                        class=""
                        type="text"
                        placeholder="Phone 2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={agenadmin_status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {agenadmin_create_field}<br />
                            Update : {agenadmin_update_field}
                        </div>
                    </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleMasterAgenAdminSave();
            }} 
            button_function="SAVE"
            button_title="Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>