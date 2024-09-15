document.addEventListener('DOMContentLoaded', function() {
    fetch('/podcasts')
        .then(response => response.json())
        .then(data => {
            const podcastList = document.getElementById('podcast-list');
            data.forEach(podcast => {
                const li = document.createElement('li');
                li.textContent = podcast.title + " - " + podcast.description;
                podcastList.appendChild(li);
            });
        });

    fetch('/recommendations?user_id=1')
        .then(response => response.json())
        .then(data => {
            const recommendationList = document.getElementById('recommendation-list');
            data.forEach(recommendation => {
                const li = document.createElement('li');
                li.textContent = recommendation.title;
                recommendationList.appendChild(li);
            });
        });
});
