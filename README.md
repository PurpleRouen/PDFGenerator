# PDFGenerator (Génération et remplissage de PDF)
Ce microservice développé en Golang permet d'utiliser un fichier PDF en tant que template pour générer un nouveau fichier PDF en y insérant des données.
Le payload (données en entrées) ainsi que les positions des données à insérer sont définis dans le fichier `watermarkData.go`.
Le fichier `watermarkData.go` est un exemple de données à insérer dans le PDF.

Exemple d'intégration avec PHP (Symfony) :
```php
public function getRegistrationPdf(): Response
{
    // ...
    $response = $httpClient->request('POST', 'http://127.0.0.1:8080/generate-pdf', [
        'json' => [
            // your payload (c.f. watermarkData.go)
        ],
    ]);

    $fileContent = $response->getContent();
    // ...
}
```
