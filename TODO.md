## TODO

1. Il bastionn host deve autenticarsi al vault (auth-plugin/admin-login), questo path non necessita di autorizzazioni. per ricevere un token che gli permetta di verificare il contesto utente (autenticare l'utente con l'api sign in) 1 token viene generato dal vault per chiamare le altre funzioni del plugin e l'altro è un JWT applicativo del server per permettere successive richieste

2. Tabella admin con credenziali del bastion host: autentico il bastion host con l'API, creo un jwt che viene salvato in memoria del vault e mi servirà per per le le altre operazioni (login utente, registrazione utente, get utene/i) in questo modo quando  il BH è autenticato riceve un token dal vault col quale può usare le altre funzioni del plugin 

3. Tabella Log in cui salvo tutte le richieste che partono dal vault (signin/signup...) Qua ho bisogno del jwt del BH per fare la get dei records
Sostanzialmente tutte gli endopoints hanno bisogno del JWT che viene assegnato solo al BH per autenticare le richieste 