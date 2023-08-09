<script>
    import Home from "../adminrule/Home.svelte";
    import Entry from "../adminrule/Entry.svelte";

    let listAdminrule = [];
    let sData = "";
    let record = "";
    let totalrecord = 0;
    let adminrule_idadmin = "";
    let adminrule_rule = "";
    export let table_header_font = "";
    export let table_body_font = "";
    let token = localStorage.getItem("token");
    let akses_page = false;

    async function initapp() {
        const res = await fetch("/api/valid", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                page: "ADMINRULE-VIEW",
            }),
        });
        const json = await res.json();
        if (json.status === 400) {
            logout();
        } else if (json.status == 403) {
            alert(json.message);
        } else {
            akses_page = true;
            initAdminrule();
        }
    }
    async function initAdminrule() {
        listAdminrule = [];
        const res = await fetch("/api/alladminrule", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({}),
        });
        const json = await res.json();
        if (json.status == 200) {
            record = json.record;
            totalrecord = record.length;
            let no = 0;
            if (record != null) {
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listAdminrule = [
                        ...listAdminrule,
                        {
                            adminrule_no: no,
                            adminrule_idadmin: record[i]["adminrule_idadmin"],
                            adminrule_rule: record[i]["adminrule_rule"],
                        },
                    ];
                }
            }
        } else {
            logout();
        }
    }
    async function logout() {
        localStorage.clear();
        window.location.href = "/";
    }
    const handleBackHalaman = () => {
        adminrule_idadmin = "";
        adminrule_rule = "";
        sData = "";
        listAdminrule = [];
        handleRefreshData("all");
    };
    const handleEditData = (e) => {
        adminrule_idadmin = e.detail.e;
        adminrule_rule = e.detail.f;
        sData = "Edit";
    };
    const handleRefreshData = (e) => {
        listAdminrule = [];
        totalrecord = 0;
        setTimeout(function () {
            initAdminrule();
        }, 500);
    };
    initapp();
</script>

{#if akses_page == true}
    {#if sData == "" && adminrule_idadmin == ""}
        <Home
            on:handleEditData={handleEditData}
            on:handleRefreshData={handleRefreshData}
            {token}
            {table_header_font}
            {table_body_font}
            {listAdminrule}
            {totalrecord}
        />
    {:else}
        <Entry
            on:handleBackHalaman={handleBackHalaman}
            {sData}
            {token}
            {table_header_font}
            {table_body_font}
            {adminrule_idadmin}
            {adminrule_rule}
        />
    {/if}
{/if}
