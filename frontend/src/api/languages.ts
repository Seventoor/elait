import { config } from "@/config"

interface Language {
    name: string
    code: string
}

export const fetchLanguageList = async (): Promise<Language[]> => {
    const response = await fetch(`${config.backendURL}/lang-list`)

    if (!response.ok) {
        throw new Error('Failed to fetch languages')
    }

    return response.json()
}