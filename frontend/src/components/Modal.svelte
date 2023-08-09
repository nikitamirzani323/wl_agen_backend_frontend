<script>
    export let modal_id = "modal_1"
    export let modal_size = "modal-dialog-centered"
    export let modal_title = "Title"
    export let modal_body_css = ""
    export let modal_footer_css = ""
    export let modal_footer = false
    export let modal_search = false
</script>

<div class="modal fade" id={modal_id} tabindex="-1" data-bs-backdrop="false">
	<div class="modal-dialog {modal_size}">
		<div class="modal-content" style="border-radius: 0px;">
			<div class="modal-header">
				<h5 class="modal-title" id="exampleModalLabel">{modal_title}</h5>
				<button
					type="button"
					class="btn-close"
					data-bs-dismiss="modal"
					aria-label="Close"></button>
			</div>
			{#if modal_search}
				<slot name="search" />
			{/if}
			<div class="modal-body" style="{modal_body_css}">
				<slot name="body" />
			</div>
            {#if modal_footer}
                <div class="modal-footer" style="{modal_footer_css}">
                    <slot name="footer" />
                </div>
            {/if}
		</div>
	</div>
</div>