<script lang="ts">
	let selectedFile: File | null = null;
	let imageUrl: string = '';
	let resizedImageUrl: string = '';
	let width: number = 300;
	let height: number = 300;
	let errorMessage: string = '';
	let originalResolution: string = '';

	function handleFileUpload(event: Event): void {
		const target = event.target as HTMLInputElement;
		const file = target.files?.[0];
		if (file) {
			const reader = new FileReader();
			reader.onload = () => {
				imageUrl = reader.result as string;
				getImageResolution(imageUrl);
			};
			reader.readAsDataURL(file);
			selectedFile = file;
		}
	}

	function getImageResolution(imageSrc: string): void {
		const img = new Image();
		img.onload = () => {
			originalResolution = `${img.width}x${img.height}`;
		};
		img.src = imageSrc;
	}

	async function resizeImage(): Promise<void> {
		if (!selectedFile) return;

		errorMessage = '';

		const formData = new FormData();
		formData.append('image', selectedFile);

		try {
			const response = await fetch(`http://localhost:8080/resize?width=${width}&height=${height}`, {
				method: 'POST',
				body: formData
			});

			if (response.ok) {
				const blob = await response.blob();
				resizedImageUrl = URL.createObjectURL(blob);
				errorMessage = '';
			} else {
				const errorData = await response.json();
				errorMessage = errorData.error;
				resizedImageUrl = '';
			}
		} catch (err) {
			errorMessage = 'An unexpected error occurred. Please try again.';
			resizedImageUrl = '';
		}
	}
</script>

<main class="container text-center my-5">
	<h1 class="text-success">Image Resizer</h1>
	<div class="mb-4">
		<input type="file" accept="image/*" class="form-control-file" on:change={handleFileUpload} />
	</div>

	{#if imageUrl}
		<h2 class="text-info">Selected Image:</h2>
		<img src={imageUrl} alt="Selected" class="img-fluid mb-3" />

		<div class="text-muted mb-3">
			<p><strong>Original Resolution:</strong> {originalResolution}</p>
		</div>

		<div class="input-group justify-content-center mb-4">
			<div class="input-group-prepend">
				<label for="widthInput" class="input-group-text">Width:</label>
			</div>
			<input
				type="number"
				id="widthInput"
				bind:value={width}
				class="form-control w-25"
				aria-label="Width"
			/>

			<div class="input-group-prepend">
				<label for="heightInput" class="input-group-text">Height:</label>
			</div>
			<input
				type="number"
				id="heightInput"
				bind:value={height}
				class="form-control w-25"
				aria-label="Height"
			/>
			<div class="input-group-append">
				<button class="btn btn-success" on:click={resizeImage}>Resize Image</button>
			</div>
		</div>

		{#if errorMessage}
			<div class="alert alert-danger">{errorMessage}</div>
		{/if}
	{/if}

	{#if resizedImageUrl}
		<h2 class="text-success">Resized Image:</h2>
		<img src={resizedImageUrl} alt="Resized" class="img-fluid mb-3" />

		<a href={resizedImageUrl} download="resized_image.jpg">
			<button class="btn btn-primary">Download Resized Image</button>
		</a>
	{/if}
</main>

<style>
	img {
		border: 1px solid #ccc;
		margin-top: 10px;
		border-radius: 5px;
	}
	.input-group-prepend,
	.input-group-append {
		display: flex;
		align-items: center;
	}
</style>
