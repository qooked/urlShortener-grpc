window.onload = async (e) => {
  const formData = new FormData();
  formData.append('url', window.location.href);
  await fetch("/getOriginalURL", {
    method: "POST",
    body: formData,
  })
    .then(async (response) => {
      console.log(response.status);
      if (response.status === 200) {
        response.json().then((data) => {
          let originalURL = data.OriginalURL;
          window.location.href = originalURL;
          alert(originalURL)
        });
        return;
      }else{
        window.location.href = "/notFound";
      }
    })
    .catch((error) => {
      alert(error);
    });
};