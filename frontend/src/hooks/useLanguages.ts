import { useQuery } from "@tanstack/react-query"
import { fetchLanguageList } from "@/api/languages"

export function useLanguages() {
    return useQuery({
        queryKey: ['languages'],
        queryFn: fetchLanguageList,
    })
}